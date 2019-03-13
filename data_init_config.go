package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"gpgw/app"
	"gpgw/db/config"
	"gpgw/model"
	"gputils"
	"gputils/logger"
	"gputils/status"
	"gputils/util"
)

func main() {
	env := gputils.DefineRunTimeCommonFlag()
	logger.InitLogger(logger.LvlDebug, nil)
	logger.Info("run info=====", "env", env, "docker_env", gputils.GetUseDocker())
	config.GetGateWayDbConfig(env)
	config.GetDb()

	ModifyBankDataV4_2_1()
}

func InitConfig()  {
	// 创建默认配置表 帮助中心
	createHelpCenter()
	if key, err := genDefaultPrivateKeyInGW(); err != nil {
		fmt.Println("genDefaultPrivateKeyInGW#generate private key for PAYMAYA exists an error", "err", err.Error())
	} else {
		fmt.Println("generate private key for PAYMAYA successfully", "the value is:", key)
	}
	if err := genLoanSwitchForPeso2go(); err != nil {
		fmt.Println("genLoanSwitchForPeso2go#generate the loan switch exists an error", "err", err.Error())
	}
	if err := genLoanSwitchForCashNiJuan(); err != nil {
		fmt.Println("genLoanSwitchForCashNiJuan#generate the loan switch exists an error", "err", err.Error())
	}
	if err := GenDefaultAgreementForPeso2go(); err != nil {
		fmt.Println("genDefaultAgreementForPeso2go#generate the default agreement of the Peso2go exists an error", "err", err.Error())
	}
	if err := GenDefaultBankInPeso2goForAPP(); err != nil {
		fmt.Println("genDefaultBankInPeso2goForAPP#fail to generate the default bank in Peso2go for APP", "err", err.Error())
	}
}


//生成PAYMAYA默认的私钥
func genDefaultPrivateKeyInGW() (key string, err error) {
	conf := config.GetDb().Model(&model.PrivateKey{})
	var tem model.PrivateKey
	if searErr := conf.Where("length = ? AND channel = ?", app.PrivateKeyLength, gputils.PayChannelPaymaya).First(&tem).Error; searErr != nil && searErr != gorm.ErrRecordNotFound {
		logger.Error("getDefaultPrivateKeyInGW#search the private key record exists an error", "err", searErr.Error())
		err = errors.New(status.GwSystemException)
		return
	} else if searErr == gorm.ErrRecordNotFound {
		var temp model.PrivateKey
		temp.Key = util.GetCommonPrivateKey(app.PrivateKeyLength)
		temp.Length = app.PrivateKeyLength
		temp.Channel = gputils.PayChannelPaymaya
		if createErr := conf.Create(&temp).Error; createErr != nil {
			logger.Error("GetPrivateKey#create the private key record exists an error", "err", createErr.Error())
			err = errors.New(status.GwSystemException)
			return
		}
		key = temp.Key
		return
	}
	key = tem.Key
	return
}

//生成Peso2go自动放款的开关
func genLoanSwitchForPeso2go() (err error) {
	conf := config.GetDb().Model(&model.GateWayConf{})
	var tem model.GateWayConf
	if searErr := conf.Where("name = ?", app.AutoLoanSwitchForPeso2go).First(&tem).Error; searErr != nil && searErr != gorm.ErrRecordNotFound {
		logger.Error("genLoanSwitchForPeso2go#search the GateWayConf exists an error", "err", searErr.Error())
		err = errors.New(status.GwSystemException)
		return
	} else if searErr == gorm.ErrRecordNotFound {
		var temp model.GateWayConf
		temp.Name = app.AutoLoanSwitchForPeso2go
		temp.Value = app.AutoLoanSwitchOff
		temp.ConfType = app.AutoLoanSwitchType
		temp.Available = app.ConfAvailable
		temp.Remark = "the switch of the AutoLoan for Peso2go"
		if createErr := conf.Create(&temp).Error; createErr != nil {
			logger.Error("genLoanSwitchForPeso2go#create the GateWayConf exists an error", "err", createErr.Error())
			err = errors.New(status.GwSystemException)
			return
		}
	}
	return
}

//生成Peso2go自动放款的开关
func genLoanSwitchForCashNiJuan() (err error) {
	conf := config.GetDb().Model(&model.GateWayConf{})
	var tem model.GateWayConf
	if searErr := conf.Where("name = ?", app.AutoLoanSwitchForCashNiJuan).First(&tem).Error; searErr != nil && searErr != gorm.ErrRecordNotFound {
		logger.Error("genLoanSwitchForCashNiJuan#search the GateWayConf exists an error", "err", searErr.Error())
		err = errors.New(status.GwSystemException)
		return
	} else if searErr == gorm.ErrRecordNotFound {
		var temp model.GateWayConf
		temp.Name = app.AutoLoanSwitchForCashNiJuan
		temp.Value = app.AutoLoanSwitchOff
		temp.ConfType = app.AutoLoanSwitchType
		temp.Available = app.ConfAvailable
		temp.Remark = "the switch of the AutoLoan for CashNiJuan"
		if createErr := conf.Create(&temp).Error; createErr != nil {
			logger.Error("genLoanSwitchForCashNiJuan#create the GateWayConf exists an error", "err", createErr.Error())
			err = errors.New(status.GwSystemException)
			return
		}
	}
	return
}

//生成Peso2go的默认配置
func GenDefaultAgreementForPeso2go() (err error) {
	conf := config.GetDb().Model(&model.GateWayConf{})
	var tem model.GateWayConf
	if searErr := conf.Where("name = ?", app.Peso2goAgreementEnglish).First(&tem).Error; searErr != nil && searErr != gorm.ErrRecordNotFound {
		logger.Error("genDefaultAgreementForPeso2go#search the GateWayConf exists an error", "err", searErr.Error())
		err = errors.New(status.GwSystemException)
		return
	} else if searErr == gorm.ErrRecordNotFound {
		var temp model.GateWayConf
		temp.Name = app.Peso2goAgreementEnglish
		valueStr := base64.StdEncoding.EncodeToString([]byte(model.AgreementForPeso2go))
		temp.Value = valueStr
		temp.ConfType = app.Peso2goAgreementEnglish
		temp.Available = app.ConfAvailable
		temp.Remark = "the default agreement of the Peso2go which is converted to base64"
		if createErr := conf.Create(&temp).Error; createErr != nil {
			logger.Error("genDefaultAgreementForPeso2go#create the default agreement of the Peso2go exists an error", "err", createErr.Error())
			err = errors.New(status.GwSystemException)
			return
		}
	}
	return
}

//生成Peso2go的APP端的收款银行
func GenDefaultBankInPeso2goForAPP() (err error) {
	data := model.GetDefaultBankInPeso2goForAPP()
	for i := range data {
		conf := config.GetDb().Model(&model.GateWayConf{})
		var tem model.GateWayConf
		if searErr := conf.Where("value = ? AND conf_type = ?", data[i].Value, data[i].ConfType).First(&tem).Error; searErr != nil && searErr != gorm.ErrRecordNotFound {
			logger.Error("genDefaultBankInPeso2goForAPP#search the GateWayConf exists an error", "err", searErr.Error())
			err = errors.New(status.GwSystemException)
			return
		} else if searErr == gorm.ErrRecordNotFound {
			var temp model.GateWayConf
			temp.Name = data[i].Name
			temp.Value = data[i].Value
			temp.ConfType = data[i].ConfType
			temp.Available = data[i].Available
			temp.Remark = data[i].Remark
			if createErr := conf.Create(&temp).Error; createErr != nil {
				logger.Error("genDefaultBankInPeso2goForAPP#fail to create default bank in Peso2go for APP", "err", createErr.Error(), "bank record", util.StringifyJson(temp))
				err = errors.New(status.GwSystemException)
				return
			}
		}
	}
	return
}

// 默认创建帮助中心
func createHelpCenter() {
	// 默认创建
	for _, item := range model.GetDefaultHelpCenter() {
		tempVal := model.HelpCenter{
			Type:      item.Type,
			Level:     item.Level,
			Language:  item.Language,
			Title:     item.Title,
			Detail:    item.Detail,
			Number:    item.Number,
			Available: item.Available,
		}
		if err := config.GetDb().Model(&model.HelpCenter{}).FirstOrCreate(&tempVal, &tempVal).Error; err != nil {
			logger.Error("createHelpCenter#create help center error", "err", err.Error(), "tempVal", util.StringifyJson(tempVal))
		}
	}
}

// V4_2_1版本添加银行信息提示
func ModifyBankDataV4_2_1() {
	logger.Info("ModifyBankDataV4_2_1 begin ...")
	tx := config.GetDb().Begin()
	defer tx.Rollback()

	if err := tx.Model(&model.GateWayConf{}).Update("remark_available", "0").Error; err != nil {
		logger.Error("ModifyBankDataV4_2_1#modify bank info error", err, "remark_available", "0")
		tx.Rollback()
		return
	}

	if err := tx.Model(&model.GateWayConf{}).Where("Value = ?", "BDO").Updates(
		map[string]interface{}{
			"remark_available": "1",
			"remark": "Notice! Don't use BDO Cash Card as your Recipient Account!",
		}).Error; err != nil {
		logger.Error("ModifyBankDataV4_2_1#modify bank info error", "err", err, "Value", "BDO")
		tx.Rollback()
		return
	}

	if err := tx.Model(&model.GateWayConf{}).Where("Value = ?", "BPI").Updates(
		map[string]interface{}{
			"remark_available":"1",
			"remark":"Notice! Don't use BPI Family Bank account as your Recipient Account!",
		}).Error; err != nil {
		logger.Error("ModifyBankDataV4_2_1#modify bank info error", "err", err, "Value", "BPI")
		tx.Rollback()
		return
	}

	tx.Commit()
	logger.Info("ModifyBankDataV4_2_1 finished...")
}