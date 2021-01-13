from typing import Optional, Dict

import spacy
from spacy.language import Language
from spacy.tokens import Doc

from presidio_analyzer import PresidioLogger
from presidio_analyzer.nlp_engine import NlpArtifacts, NlpEngine

logger = PresidioLogger()


class SpacyNlpEngine(NlpEngine):
    """SpacyNlpEngine is an abstraction layer over the nlp module.
    It provides processing functionality as well as other queries
    on tokens.
    The SpacyNlpEngine uses SpaCy as its NLP module
    """

    engine_name = "spacy"
    is_available = bool(spacy)

    def __init__(self, models: Optional[Dict[str, str]] = None):
        """
        Initializes a wrapper on spaCy functionality.
        :param models: Dictionary with the name of the spaCy model per language.
        For example: models = {"en": "en_core_web_lg"}
        """
        if not models:
            models = {"en": "en_core_web_lg"}
        logger.debug(f"Loading SpaCy models: {models.values()}")

        self.nlp = {
            lang_code: spacy.load(model_name, disable=["parser", "tagger"])
            for lang_code, model_name in models.items()
        }

    def process_text(self, text: str, language: str):
        """Execute the SpaCy NLP pipeline on the given text
        and language
        """
        doc = self.nlp[language](text)
        return self.doc_to_nlp_artifact(doc, language)

    def is_stopword(self, word: str, language: str) -> bool:
        """returns true if the given word is a stop word
        (within the given language)
        """
        return self.nlp[language].vocab[word].is_stop

    def is_punct(self, word: str, language: str) -> bool:
        """returns true if the given word is a punctuation word
        (within the given language)
        """
        return self.nlp[language].vocab[word].is_punct

    def get_nlp(self, language: str) -> Language:
        return self.nlp[language]

    def doc_to_nlp_artifact(self, doc: Doc, language: str) -> NlpArtifacts:
        lemmas = [token.lemma_ for token in doc]
        tokens_indices = [token.idx for token in doc]
        entities = doc.ents
        return NlpArtifacts(
            entities=entities,
            tokens=doc,
            tokens_indices=tokens_indices,
            lemmas=lemmas,
            nlp_engine=self,
            language=language,
        )
