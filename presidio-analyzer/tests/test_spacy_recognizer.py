from analyzer.predefined_recognizers import SpacyRecognizer

NER_STRENGTH = 0.85
spacy_recognizer = SpacyRecognizer()
spacy_recognizer.load()
entities = ["PERSON", "DATE_TIME"]


def test_person_first_name():
    name = 'Dan'
    results = spacy_recognizer.analyze(name, entities)

    assert len(results) == 0


def test_person_first_name_with_context():
    name = 'Dan'
    context = 'my name is '
    results = spacy_recognizer.analyze(context + name, entities)

    assert len(results) == 1
    assert results[0].score >= NER_STRENGTH
    assert results[0].entity_type == entities[0]
    assert results[0].start == 11
    assert results[0].end == 14


def test_person_full_name():
    name = 'Dan Tailor'
    results = spacy_recognizer.analyze(name, entities)

    assert len(results) == 1
    assert results[0].score >= NER_STRENGTH
    assert results[0].entity_type == entities[0]
    assert results[0].start == 0
    assert results[0].end == 10


def test_person_full_name_with_context():
    name = 'John Oliver'
    results = spacy_recognizer.analyze(name + " is the funniest comedian", entities)

    assert len(results) == 1
    assert results[0].score >= NER_STRENGTH
    assert results[0].entity_type == entities[0]
    assert results[0].start == 0
    assert results[0].end == 11


def test_person_last_name():
    name = 'Tailor'
    results = spacy_recognizer.analyze(name, entities)

    assert len(results) == 0


def test_person_full_middle_name():
    name = 'Richard Milhous Nixon'
    results = spacy_recognizer.analyze(name, entities)

    assert len(results) == 1
    assert results[0].score >= NER_STRENGTH
    assert results[0].entity_type == entities[0]
    assert results[0].start == 0
    assert results[0].end == 21


def test_person_full_middle_letter_name():
    name = 'Richard M. Nixon'
    results = spacy_recognizer.analyze(name, entities)

    assert len(results) == 1
    assert results[0].score >= NER_STRENGTH
    assert results[0].entity_type == entities[0]
    assert results[0].start == 0
    assert results[0].end == 16


def test_person_full_name_complex():
    name = 'Richard (Ric) C. Henderson'
    results = spacy_recognizer.analyze(name, entities)

    assert len(results) == 1
    assert results[0].score >= NER_STRENGTH
    assert results[0].entity_type == entities[0]
    assert results[0].start == 0
    assert results[0].end == 26


def test_date_time_simple():
    name = 'May 1st'
    results = spacy_recognizer.analyze(name + " is the workers holiday", ["DATE_TIME"])

    assert len(results) == 1
    assert results[0].score >= NER_STRENGTH
    assert results[0].entity_type == entities[1]
    assert results[0].start == 0
    assert results[0].end == 7
