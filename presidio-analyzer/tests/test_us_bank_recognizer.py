import pytest

from tests import assert_result
from presidio_analyzer.predefined_recognizers import UsBankRecognizer


@pytest.fixture(scope="module")
def recognizer():
    return UsBankRecognizer()


@pytest.fixture(scope="module")
def entities():
    return ["US_BANK_NUMBER"]


@pytest.mark.parametrize(
    # fmt: off
    "text, expected_len, expected_positions, expected_score",
    [
        # valid bank accounts
        ("945456787654", 1, ((0, 12),), 0.05),
        # invalid bank accounts
        ("1234567", 0, (), -1.0),
    ],
)
def test_all_us_banks(
    text, expected_len, expected_positions, expected_score, recognizer, entities
):
    results = recognizer.analyze(text, entities)
    assert len(results) == expected_len
    for res, (st_pos, fn_pos) in zip(results, expected_positions):
        assert_result(res, entities[0], st_pos, fn_pos, expected_score)
