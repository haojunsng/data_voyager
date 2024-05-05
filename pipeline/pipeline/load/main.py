from utils.pg_helper import SupaCursor


def main():
    postgres_cursor = SupaCursor()

    query = """
    SELECT 1
    """
    postgres_cursor.supa_run(query)


if __name__ == "__main__":
    main()
