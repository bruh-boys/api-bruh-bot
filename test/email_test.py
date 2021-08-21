import requests

"""
here im just making a simple request for test

"""

r = requests.post("http://localhost:5000/email",
                  json={"email": "monda@monda.com"})
print(r.text)
