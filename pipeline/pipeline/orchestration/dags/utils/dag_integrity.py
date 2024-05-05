import unittest
from pathlib import Path
from airflow.models import DagBag


class TestDags(unittest.TestCase):

    LOAD_THRESHOLD_SECONDS = 2

    def setUp(self):
        p = Path(__file__).parents[1]
        self.dagbag = DagBag(dag_folder=f"{p}")

    def test_dags_syntax(self):
        """Assert DAG bag load correctly"""
        for key in self.dagbag.dags:
            print(key)
        self.assertTrue(len(self.dagbag.dags) > 0, "need at least 1 dag")
        self.assertFalse(
            len(self.dagbag.import_errors),
            f"DAG import errors. Errors: {self.dagbag.import_errors}",
        )


if __name__ == "__main__":
    SUITE = unittest.TestLoader().loadTestsFromTestCase(TestDags)
    unittest.TextTestRunner(verbosity=2).run(SUITE)
