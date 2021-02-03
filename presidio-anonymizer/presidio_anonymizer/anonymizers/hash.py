"""Hashes the PII text entity."""
from hashlib import sha256, sha512, md5
from presidio_anonymizer.anonymizers import Anonymizer
from presidio_anonymizer.anonymizers.validators import validate_parameter_in_range


class Hash(Anonymizer):
    """Hash given text with sha256/sha512/md5 algorithm."""

    HASH_TYPE = "hash_type"
    SHA256 = "sha256"
    SHA512 = "sha512"
    MD5 = "md5"

    def anonymize(self, text: str = None, params: dict = None) -> str:
        """
        Hash given value using sha256.

        :return: hashed original text
        """
        hash_type = params.get(self.HASH_TYPE, self.SHA256)
        hash_switcher = {
            self.SHA256: lambda s: sha256(s.encode()),
            self.SHA512: lambda s: sha512(s.encode()),
            self.MD5: lambda s: md5(s.encode()),
        }
        return hash_switcher.get(hash_type)(text).hexdigest()

    def validate(self, params: dict = None) -> None:
        """Validate the hash type is string and in range of allowed hash types."""
        validate_parameter_in_range(
            [self.SHA256, self.SHA512, self.MD5],
            params.get(self.HASH_TYPE, self.SHA256),
            self.HASH_TYPE,
            str,
        )
        pass
