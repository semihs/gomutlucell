# GoMutlucell ( Mutlucell golang kütüphanesi www.mutlucell.com )

## Özellikler
* Sms Gönderim

## Kurulum
```
go get github.com/semihs/goform
```

## Kullanım

### Sms Gönderim

```go
package main

import (
	"github.com/semihs/gomutlucell"
	"fmt"
)


func main() {
	mutluCellClient := gomutlucell.NewMutluCellClient("kullanıcı adınız", "şifreniz", "sms başlığı", "karakter kodlaması (türkçe için turkish)")
	if err := mutluCellClient.SendSms(gomutlucell.Message{Message: "Merhaba. Bu bir deneme mesajidir.",Numbers:"532 567 89 90, 05556667788, +905551112233, 905324441122",}); err != nil {
            fmt.Errorf("an error occurred while sms sending %s", err)
        }
}
```
