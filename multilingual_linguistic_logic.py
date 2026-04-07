def analyze_multilingual_text(text, language):
    """Simple parser to identify specific structural markers across languages."""
    words = text.lower().split()
    results = {"language": language, "word_count": len(words), "markers_found": []}

    if language == "german":
        # Passive voice and Nominalization markers
        markers = ["wird", "werden", "wurde", "ung", "keit", "heit"]
    elif language == "english":
        # Continuous forms and Passive markers
        markers = ["ing", "been", "was", "were", "by"]
    elif language == "russian":
        # Case endings and Reflexive verbs (ся/сь)
        markers = ["ся", "сь", "ого", "ему", "ыми"]
    elif language == "kazakh":
        # Agglutinative suffixes (Possessive, Plural, Case)
        markers = ["лар", "лер", "дар", "дер", "тын", "тін", "ның", "нің"]
    else:
        return "Language not supported"

    # Identify if words contain the specified structural markers
    for word in words:
        for marker in markers:
            if marker in word:
                results["markers_found"].append(word)
                break
    
    return results
