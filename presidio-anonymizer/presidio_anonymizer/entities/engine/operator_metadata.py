import logging
from abc import ABC
from typing import Dict

from presidio_anonymizer.entities import InvalidParamException
from presidio_anonymizer.operators import OperatorType, operator


class OperatorMetadata(ABC):
    """Abstract class to hold the data of the required operator."""

    def __init__(
            self,
            operator_type: OperatorType,
            params: Dict,
            operator_name: str
    ):
        self.logger = logging.getLogger("presidio-anonymizer")
        self.operator_type = operator_type
        self.operator_name = operator_name
        self.params = params
        self.__validate_fields()

    def __validate_fields(self):
        if self.operator_name is None:
            self.__validate_field("operator_name")
        if self.operator_type is None:
            self.__validate_field("operator_type")
        if self.operator_type not in operator.types:
            raise InvalidParamException(
                f"Invalid input, invalid operator type {self.operator_type}"
            )

    def __validate_field(self, field_name: str):
        self.logger.debug(f"invalid parameter, {field_name} cannot be empty")
        raise InvalidParamException(
            f"Invalid input, config must contain {field_name}"
        )