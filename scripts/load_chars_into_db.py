import os
import json
# pip install python-dotenv
from dotenv import load_dotenv
# pip install mysql-connector-python 
import mysql.connector

if __name__=="__main__":
    load_dotenv()

    config = {
        "user": "root",
        "password": os.getenv("DB_PASS"),
        "host": "localhost",
        "database": os.getenv("DB_NAME"),
        "raise_on_warnings": True
    }

    conn = mysql.connector.connect(**config)
    cur = conn.cursor()

    with open("./assets/char_jsons/1_Hayashi.json") as fh:
        file = json.load(fh)

    file["skin_path"] = "./assets/original/" + file["name"] + ".png"

    cur.execute("""
        insert into characters
        (id, name, nickname, description, rarity, skin_path) 
        values
        (%(id)s, %(name)s, %(nickname)s, %(description)s, %(rarity)s, %(skin_path)s)
    """, file)

    conn.commit()

    cur.close()
    conn.close()
