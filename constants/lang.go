package constants

type LanguageType string

const (
	LangRU LanguageType = "ru"
	LangUA LanguageType = "ua"
	LangBE LanguageType = "be"
	LangEN LanguageType = "en"
	LangES LanguageType = "es"
	LangFI LanguageType = "fi"
	LangDE LanguageType = "de"
	LangIT LanguageType = "it"
)

type LanguageTranslationType string

const (
	LangTranslationRU LanguageTranslationType = "ru"
	LangTranslationEN LanguageTranslationType = "en"
	LangTranslationES LanguageTranslationType = "es"
	LangTranslationPT LanguageTranslationType = "pt"
)

type LanguageCode int

const (
	LangCodeRU               LanguageCode = 0   // Русский
	LangCodeUK               LanguageCode = 1   // Українська
	LangCodeBE               LanguageCode = 2   // Беларуская (тарашкевiца)
	LangCodeEN               LanguageCode = 3   // English
	LangCodeES               LanguageCode = 4   // Español
	LangCodeFI               LanguageCode = 5   // Suomi
	LangCodeDE               LanguageCode = 6   // Deutsch
	LangCodeIT               LanguageCode = 7   // Italiano
	LangCodeBG               LanguageCode = 8   // Български
	LangCodeHR               LanguageCode = 9   // Hrvatski
	LangCodeHU               LanguageCode = 10  // Magyar
	LangCodeSR               LanguageCode = 11  // Српски
	LangCodePT               LanguageCode = 12  // Português
	LangCodeEL               LanguageCode = 14  // Ελληνικά
	LangCodePL               LanguageCode = 15  // Polski
	LangCodeFR               LanguageCode = 16  // Français
	LangCodeKO               LanguageCode = 17  // 한국어
	LangCodeZH               LanguageCode = 18  // 汉语
	LangCodeLT               LanguageCode = 19  // Lietuvių
	LangCodeJA               LanguageCode = 20  // 日本語
	LangCodeCS               LanguageCode = 21  // Čeština
	LangCodeET               LanguageCode = 22  // Eesti
	LangCodeTT               LanguageCode = 50  // Татарча
	LangCodeBA               LanguageCode = 51  // Башҡортса
	LangCodeCV               LanguageCode = 52  // Чăвашла
	LangCodeSK               LanguageCode = 53  // Slovenčina
	LangCodeRO               LanguageCode = 54  // Română
	LangCodeNO               LanguageCode = 55  // Norsk
	LangCodeLV               LanguageCode = 56  // Latviešu
	LangCodeAZ               LanguageCode = 57  // Azərbaycan dili
	LangCodeHY               LanguageCode = 58  // Հայերեն
	LangCodeSQ               LanguageCode = 59  // Shqip
	LangCodeSV               LanguageCode = 60  // Svenska
	LangCodeNL               LanguageCode = 61  // Nederlands
	LangCodeTK               LanguageCode = 62  // Türkmen
	LangCodeKA               LanguageCode = 63  // ქართული
	LangCodeDA               LanguageCode = 64  // Dansk
	LangCodeUZ               LanguageCode = 65  // O‘zbek
	LangCodeMO               LanguageCode = 66  // Moldovenească
	LangCodeBUA              LanguageCode = 67  // Буряад
	LangCodeTH               LanguageCode = 68  // ภาษาไทย
	LangCodeID               LanguageCode = 69  // Bahasa Indonesia
	LangCodeTG               LanguageCode = 70  // Тоҷикӣ
	LangCodeSL               LanguageCode = 71  // Slovenščina
	LangCodeBS               LanguageCode = 72  // Bosanski
	LangCodePTBR             LanguageCode = 73  // Português brasileiro
	LangCodeFA               LanguageCode = 74  // فارسی
	LangCodeVI               LanguageCode = 75  // Tiếng Việt
	LangCodeHI               LanguageCode = 76  // हिन्दी
	LangCodeSI               LanguageCode = 77  // සිංහල
	LangCodeBN               LanguageCode = 78  // বাংলা
	LangCodeTL               LanguageCode = 79  // Tagalog
	LangCodeMN               LanguageCode = 80  // Монгол
	LangCodeMY               LanguageCode = 81  // ဗမာစာ
	LangCodeTR               LanguageCode = 82  // Türkçe
	LangCodeNE               LanguageCode = 83  // नेपाली
	LangCodeUR               LanguageCode = 85  // اردو
	LangCodeKY               LanguageCode = 87  // Кыргыз тили
	LangCodePA               LanguageCode = 90  // پنجابی
	LangCodeOS               LanguageCode = 91  // Ирон
	LangCodeKN               LanguageCode = 94  // ಕನ್ನಡ
	LangCodeSW               LanguageCode = 95  // Kiswahili
	LangCodeKK               LanguageCode = 97  // Қазақша
	LangCodeAR               LanguageCode = 98  // العربية
	LangCodeHE               LanguageCode = 99  // עברית
	LangCodePreRevolutionary LanguageCode = 100 // Дореволюцiонный
	LangCodeMYV              LanguageCode = 101 // Эрзянь кель
	LangCodeKDB              LanguageCode = 102 // Адыгэбзэ
	LangCodeSAH              LanguageCode = 105 // Саха тыла
	LangCodeADY              LanguageCode = 106 // Адыгабзэ
	LangCodeUDM              LanguageCode = 107 // Удмурт
	LangCodeCHM              LanguageCode = 108 // Марий йылме
	LangCodeBE2              LanguageCode = 114 // Беларуская
	LangCodeLEZ              LanguageCode = 118 // Лезги чІал
	LangCodeTW               LanguageCode = 119 // 臺灣話
	LangCodeKUM              LanguageCode = 236 // Къумукъ тил
	LangCodeMVL              LanguageCode = 270 // Mirandés
	LangCodeSLA              LanguageCode = 298 // Русинськый
	LangCodeKRL              LanguageCode = 379 // Karjalan kieli
	LangCodeTYV              LanguageCode = 344 // Тыва дыл
	LangCodeXAL              LanguageCode = 357 // Хальмг келн
	LangCodeTLY              LanguageCode = 373 // Tolışə zıvon
	LangCodeKV               LanguageCode = 375 // Коми кыв
	LangCodeUKClassic        LanguageCode = 452 // Українська (клясична)
	LangCodeUKGalitska       LanguageCode = 454 // Українська (Галицка)
	LangCodeKAB              LanguageCode = 457 // Taqbaylit
	LangCodeEO               LanguageCode = 555 // Esperanto
	LangCodeLA               LanguageCode = 666 // Lingua Latina
)
