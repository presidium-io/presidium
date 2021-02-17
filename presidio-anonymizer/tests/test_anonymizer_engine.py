from presidio_anonymizer import AnonymizerEngine
from presidio_anonymizer.entities import AnonymizerConfig


def test_given_request_anonymizers_return_list():
    engine = AnonymizerEngine()
    expected_list = ["fpe", "hash", "mask", "redact", "replace"]
    anon_list = engine.anonymizers()

    assert anon_list == expected_list


# TODO SHIRANR enable this tests next commit.
def given_anonymizers_config_then_we_get_correct_anonymizer_with_default():
    engine = AnonymizerEngine()
    phone_number_config = AnonymizerConfig("fpe", {})
    default_config = AnonymizerConfig("redact", {})
    anonymizers = {"PHONE_NUMBER": phone_number_config,
                   "DEFAULT": default_config,
                   "NUMBER": default_config,
                   "PHONE_NUM": default_config}
    assert engine.__get_anonymizer_by_entity_type(anonymizers,
                                                  "PHONE_NUMBER") == phone_number_config
    assert engine.__get_anonymizer_by_entity_type(anonymizers,
                                                  "NONE_EXISTING") == default_config


def given_anonymizers_config_then_we_get_correct_anonymizer_without_default():
    engine = AnonymizerEngine()
    phone_number_config = AnonymizerConfig("fpe", {})
    anonymizers = {"PHONE_NUMBER": phone_number_config}
    assert engine.__get_anonymizer_by_entity_type(anonymizers,
                                                  "PHONE_NUMBER") == phone_number_config
    assert engine.__get_anonymizer_by_entity_type(anonymizers,
                                                  "NONE_EXISTING") == AnonymizerConfig(
        "replace", {})
