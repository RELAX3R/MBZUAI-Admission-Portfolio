import math

def generate_sine_wave(frequency, sample_rate=44100, duration=1.0):
    """Simulates a basic sine wave generation logic used in synthesis."""
    wave = []
    for i in range(int(sample_rate * duration)):
        # Formula: A * sin(2 * pi * f * t)
        value = math.sin(2 * math.pi * frequency * (i / sample_rate))
        wave.append(round(value, 4))
    return wave[:10]  # Return first 10 samples for demonstration

print("First 10 samples of 440Hz Sine Wave:", generate_sine_wave(440))
