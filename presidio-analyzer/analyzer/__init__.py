import os
import sys
from analyzer.pattern import Pattern  # noqa: F401
from analyzer.entity_recognizer import EntityRecognizer  # noqa: F401
from analyzer.local_recognizer import LocalRecognizer  # noqa: F401
from analyzer.recognizer_result import RecognizerResult  # noqa: F401
from analyzer.pattern_recognizer import PatternRecognizer  # noqa: F401
from analyzer.remote_recognizer import RemoteRecognizer  # noqa: F401
from analyzer.recognizer_registry.recognizer_registry import RecognizerRegistry  # noqa

sys.path.append(os.path.dirname(os.path.dirname(
    os.path.abspath(__file__))) + "/analyzer")

from analyzer.analyzer_engine import AnalyzerEngine  # noqa: F401
