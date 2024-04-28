import psycopg2
from utils.env_helper import load_env_vars


class SupaCursor:
    def __init__(self):
        self.connection_string = load_env_vars()["SUPABASE_CONNECTION_ID"]
        self.conn = None
        self.cursor = None

    def connect(self):
        try:
            self.conn = psycopg2.connect(self.connection_string)
            self.cursor = self.conn.cursor()
            print("Connected to Supabase!")
        except psycopg2.Error as e:
            print("Error: Unable to connect to the Supabase...")
            print(e)

    def execute(self, query, params=None):
        try:
            self.cursor.execute(query, params)
            print("Query executed successfully!")
        except psycopg2.Error as e:
            print("Error: Unable to execute query...")
            print(e)

    def commit(self):
        try:
            self.conn.commit()
            print("Transaction committed successfully!")
        except psycopg2.Error as e:
            print("Error: Unable to commit the transaction...")
            print(e)

    def close(self):
        if self.cursor:
            self.cursor.close()
        if self.conn:
            self.conn.close()
        print("Connection to PostgreSQL database closed")

    def supa_run(self, query):
        self.connect()
        self.execute(query)
        self.commit()
        self.close()
