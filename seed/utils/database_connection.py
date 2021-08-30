from pymongo import MongoClient, database
import os


def connect_mongo() -> MongoClient:
    mongo_uri = os.getenv("MONGO_URL")
    client = MongoClient(mongo_uri)
    return client


def get_db(client: MongoClient, name: str) -> database.Database:
    return client.get_database(name)


def get_collection(db: database.Database, name: str) -> database.Collection:
    return db.get_collection(name)
