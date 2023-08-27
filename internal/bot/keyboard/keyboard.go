package keyboard

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

// TODO удалить клавиатуру
var TestKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Product 1", "product_1"),
		tgbotapi.NewInlineKeyboardButtonData("Product 2", "product_2"),
	),
)

var MainKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("VPN", "VPN"),
	),

	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Мои ключи", "my_keys"),
	),

	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Скачать приложение", "apps"),
	),

	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("Информация", "https://duckvpn.org"),
	),

	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("Поделиться", "https://t.me/share/url?url=https://t.me/duckvpnbot?start=shared"), //TODO подумать над ссылкой
	),
)

// VPN
var VpnKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Бесплатный VPN", "free_VPN"),
	),

	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Купить подписку", "paid_VPN"),
	),

	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Назад", "main"),
	),
)

// TODO сделать динамическую клавиатуру
var FreeVPNKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🇳🇱Нидерланды", "get_free_VPN_key_NL"),
	),

	//TODO добавить стран
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🇷🇺Россия", "get_free_VPN_key_RU"),
	),

	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Назад", "VPN"),
	),
)

// TODO сделать динамическую клавиатуру
var PaidVPNKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🇳🇱Нидерланды", "get_paid_VPN_key_NL"),
	),

	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🇷🇺Россия", "get_paid_VPN_key_RU"),
	),

	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Назад", "VPN"),
	),
)

var PaidVPNNLKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("30 дней за 200₽", "get_paid_VPN_key_NL_30"),
	),

	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("120 дней за 800₽", "get_paid_VPN_key_NL_120"),
	),

	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Назад", "get_paid_VPN_key_NL"),
	),
)

var ReturnPaidVPNKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Назад", "paid_VPN"),
	),

	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Главное меню", "main"),
	),
)

var ReturnPaidKeyVPNKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Назад", "paid_VPN"),
	),

	//tgbotapi.NewInlineKeyboardRow(
	//	tgbotapi.NewInlineKeyboardButtonData("Главное меню", "main"),
	//),
)

var ReturnFreeVPNKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Назад", "free_VPN"),
	),

	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Главное меню", "main"),
	),
)

var ReturnToMain = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Главное меню", "main"),
	),
)

// KEYS
// MY KEYS
var MyKeysKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Мои платные ключи", "my_paid_keys"),
	),

	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Мои бесплатный ключ", "my_free_key"),
	),

	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Удалить мой бесплатный ключ", "delete_my_free_key"),
	),

	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Назад", "main"),
	),
)

var MyPaidKeysKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Купить ключ", "my_paid_VPN"), //точно ли paid???
	),

	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Назад", "my_keys"),
	),

	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Главное меню", "main"),
	),
)

var MyFreeKeyNotExistKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Получить бесплатный ключ", "free_VPN"),
	),

	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Назад", "my_keys"),
	),

	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Главное меню", "main"),
	),
)

var MyFreeKeyExistKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Назад", "my_keys"),
	),

	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Главное меню", "main"),
	),
)

var FreeVPNKeyKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Назад", "free_VPN"),
	),

	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Главное меню", "main"),
	),
)

// APPS
var AppsKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("iOS 🍏", "https://apps.apple.com/app/outline-app/id1356177741"),
		tgbotapi.NewInlineKeyboardButtonURL("iOS 🍎", "http://apps.apple.com/app/shadowrocket/id932747118"),
	),

	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("Android 🤖", "https://play.google.com/store/apps/details?id=org.outline.android.client"),
		tgbotapi.NewInlineKeyboardButtonURL("Android 🤖🤖", "https://s3.amazonaws.com/outline-releases/client/android/stable/Outline-Client.apk"),
	),

	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("Windows", "https://s3.amazonaws.com/outline-releases/client/windows/stable/Outline-Client.exe"),
		tgbotapi.NewInlineKeyboardButtonURL("macOS", "https://itunes.apple.com/app/outline-app/id1356178125"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Назад", "main"),
	),
)
