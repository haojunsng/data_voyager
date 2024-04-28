from utils.pg_helper import SupaCursor


def main():
    postgres_cursor = SupaCursor()

    query = """

    """
    postgres_cursor.supa_run(query)


if __name__ == "__main__":
    main()
