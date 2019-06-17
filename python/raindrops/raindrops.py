def raindrops(number):
    raindrop_speak = ""

    if number % 3 == 0:
        raindrop_speak += "Pling"
    if number % 5 == 0:
        raindrop_speak += "Plang"
    if number % 7 == 0:
        raindrop_speak += "Plong"

    if raindrop_speak == "":
        return str(number)

    return raindrop_speak
