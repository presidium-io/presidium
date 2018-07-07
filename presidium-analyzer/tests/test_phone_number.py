from analyzer import matcher
from analyzer import common_pb2
import logging
import os

fieldType = common_pb2.FieldTypes()
fieldType.name = "PHONE_NUMBER"
types = [fieldType]


def test_phone_number_simple():
    match = matcher.Matcher()
    number = '052-5552606'
    results = match.analyze_text('my phone number is ' + number, types)

    assert results[0].text == number


def test_phone_number_text1():
    path = os.path.dirname(__file__) + '/data/text1.txt'
    text_file = open(path, 'r')
    match = matcher.Matcher()
    results = match.analyze_text(text_file.read(), types)
    assert (len(results) == 2)
