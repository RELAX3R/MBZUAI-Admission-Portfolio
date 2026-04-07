def analyze_german_structure(text):
    """Simple parser to identify potential passive voice markers in German."""
    # Markers for Passive: 'wird', 'werden', 'wurde' + 'ge-' prefix in participle
    passive_markers = ["wird", "werden", "wurde", "worden"]
    words = text.lower().split()
    
    findings = [word for word in words if word in passive_markers]
    
    return {
        "word_count": len(words),
        "passive_indicators_found": len(findings),
        "markers": findings
    }
