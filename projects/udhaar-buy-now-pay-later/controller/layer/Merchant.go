package layer

import (
	"fmt"
	"udhaar/constants"
	"udhaar/models"
	"udhaar/utils"

	"github.com/spf13/cast"
)

/*
new merchant m1 2%
*/

func AddMerchant(command []string) (newMerchantDetails string, err error) {

	if len(command) != 4 {
		err = constants.ErrInvalidCommand
		return
	}

	name, err := utils.IsValidName(command[2])

	if err != nil {
		return
	}

	discount, err := utils.IsValidDiscount(command[3])

	if err != nil {
		return
	}

	// if the user is already present or not

	_, ok := models.MerchantMap[name]
	// if the user is found htne give MerchantAlreadyExistError
	if ok {
		err = constants.ErrMerchantAlreadyExists
		return
	}

	newMerchant := models.Merchant{
		UniqueName:      name,
		DiscountOffered: discount,
	}

	models.MerchantMap[name] = newMerchant

	newMerchantDetails = fmt.Sprintf(
		"%v ( %v )",
		newMerchant.UniqueName,
		newMerchant.DiscountOffered,
	)

	return
}

func UpdateMerchant(command []string) (newDisountString string, err error) {

	/*
		update merchant m1 1%
	*/

	if len(command) != 4 {
		err = constants.ErrInvalidCommand
		return
	}

	name := command[2]

	// check if the merchant exists
	merchantObj, isExists := models.MerchantMap[name]
	if !isExists {
		err = constants.ErrMerchantMissing
		return
	}

	newDiscountFloat, err := utils.IsValidDiscount(command[3])
	if err != nil {
		return
	}

	merchantObj.DiscountOffered = newDiscountFloat
	models.MerchantMap[name] = merchantObj
	newDisountString = cast.ToString(newDiscountFloat)
	return
}

func GetDiscount(command []string) (discountString string, err error) {

	/*
		report discount m1
	*/
	if len(command) != 3 {
		err = constants.ErrInvalidCommand
		return
	}

	merchantName := command[2]

	// check if the user/merchant exists or not

	merchantObject, isUserExist := models.MerchantMap[merchantName]

	// if the user doesn't exist
	if !isUserExist {
		err = constants.ErrMerchantMissing
		return
	}
	discountString = cast.ToString(merchantObject.DiscountOffered)

	return

}
