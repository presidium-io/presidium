from entity_recognizer import EntityRecognizer 
from recognizer_result import RecognizerResult
from pattern import Pattern
import datetime
import re2 as re


class PatternRecognizer(EntityRecognizer):

    def __init__(self, supported_entities, supported_languages, patterns=None,
                 black_list=None, context=None, version="0.0.1"):
        """
            :param patterns: the list of patterns to detect
            :param black_list: the list of words to detect
            :param context: list of context words
        """
        if len(supported_entities) > 1 and len(patterns) > 0:
            raise ValueError("Pattern recognizer supports only one entity")

        super().__init__(supported_entities, supported_languages, version)
        if patterns is None:
            self.patterns = []
        else:
            self.patterns = patterns
        self.context = context

        if black_list:
            black_list_pattern = PatternRecognizer.__black_list_to_regex(black_list)
            self.patterns.append(black_list_pattern)

    def load(self):
        pass

    def analyze_all(self, text, entities):
        """

        :param text: the text to analyze
        :param entities: the entities to be detected
        :return: the analyze result- the detected entities and their location in the text
        """
        results = []

        if len(self.patterns) > 0:
            pattern_result = self.__analyze_regex_patterns(text)

            if pattern_result:
                results.extend(pattern_result)

        analyzed_text = self.analyze_text(text, entities)
        if analyzed_text:
            results.extend(analyzed_text)

        return results

    @staticmethod
    def __black_list_to_regex(black_list):
        """
          Convert a list of word to matching regex
        :param black_list: the list of words to detect
        :return:the regex of the words for detection
        """
        regex = r"(?:^|(?<= ))(" + '|'.join(black_list) + r")(?:(?= )|$)"
        return Pattern("black_list", 1.0, regex)

    def validate_pattern_logic(self, pattern_text, result):
        """
        Validate the pattern logic, for example checksum method.
        :param text: the pattern text to validate
        :param result: the current result of the pattern (before
        :return: The updated result - if the pattern is valid the score can be raised
        """
        return result


    def __analyze_regex_patterns(self, input_text):
        """Check for specific pattern in text

        Args:
          :return: the detection results
          :param input_text:text to analyze
        """
        results = []
        for pattern in self.patterns:
            match_start_time = datetime.datetime.now()
            matches = re.finditer(
                pattern.pattern,
                input_text,
                flags=re.IGNORECASE | re.DOTALL | re.MULTILINE)
            match_time = datetime.datetime.now() - match_start_time
            self.logger.debug('--- match_time[{}]: {}.{} seconds'.format(
                pattern.name, match_time.seconds, match_time.microseconds))

            for match in matches:
                start, end = match.span()
                current_match = input_text[start:end]

                # Skip empty results
                if current_match == '':
                    continue

                res = RecognizerResult(start, end, pattern.strength, self.supported_entities[0])
                res = self.validate_pattern_logic(current_match, res)

                if res is None or res.score == 0:
                    continue

                results.append(res)

        return results

