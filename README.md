<h1>a simple api that sent spam via sms and email</h1>

routes:

- /sms

- /email


example request with python (sms):

```py
import requests

r = requests.post("https://api-bruh-bot.elpanajose.repl.co/sms",
                  json={"phone": "+5759075894789"})
print(r.text)
```

email:

```py
import requests

r = requests.post("https://api-bruh-bot.elpanajose.repl.co/email",
                  json={"email": "monda@monda.com"})
print(r.text)
```




