<h1>a simple api that sent spam via sms and email</h1>

routes:

- /sms

- /email (Its not working beacuse the Host, but you can use it by yourself)


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

<h1>Do you want to use it locally?</h1>

- install al the go dependencies with ```go get .```

- Then install all the python deps with ```pip3 install -r requirements.txt```

- and finally in the ```/routes``` directory create a .env file with the following content

```
EMAIL1=
EMAIL2=
PASSWORD=
```
