input_file = open('input.txt', 'r')
input_lines = input_file.readlines()

hand_count = len(input_lines)

character_priority_map = {
    "A": "m",
    "K": "l",
    "Q": "k",
    "T": "j",
    "9": "i",
    "8": "h",
    "7": "g",
    "6": "f",
    "5": "e",
    "4": "d",
    "3": "c",
    "2": "b",
    "J": "a"
}
hand_type_priority_map = {
    "5": 700,
    "41":600,
    "32": 500,
    "311": 400,
    "221": 300,
    "2111": 200,
    "11111": 100
}
input_map = {}

def generate_hand_type(character_count):
    max_count = -1
    max_count_character = ""
    counts = []
    for key, value in character_count.items():
        if value != -1:
            if key != "J" and max_count < value:
                max_count = value
                max_count_character = key
            counts.append(value)
    counts.sort(reverse=True)
    hand_type = "".join([str(c) for c in counts])

    if max_count_character == "":
        max_count_character = "J"

    return hand_type, max_count_character

def detect_hand_type(hand):
    character_count = {}
    for char in hand:
        if char in character_count:
            character_count[char] += 1
        else:
            character_count[char] = 1
    
    hand_type, max_count_character = generate_hand_type(character_count)

    if max_count_character == "J":
        return hand_type, hand_type
    alternative_hand_type = ""

    if "J" in character_count:
        character_count[max_count_character] += character_count["J"]
        character_count["J"] = -1
        alternative_hand_type, _ = generate_hand_type(character_count)
    else:
        alternative_hand_type = hand_type

    return hand_type, alternative_hand_type

def get_hand_value_by_character_priority(hand):
    output = ""
    for char in hand:
        output += character_priority_map[char]

    return output

if __name__ == "__main__":
    for line in input_lines:
        line = line.split()
        hand = line[0]
        bid = int(line[1])
        hand_type, alternative_hand_type = detect_hand_type(hand)
        hand_value_by_character_priority = get_hand_value_by_character_priority(hand)
        input_map[hand] = {"hand_type_value": hand_type_priority_map[alternative_hand_type], "hand_value_by_character_priority": hand_value_by_character_priority, "bid": bid}

    sorted_by_character_priority = dict(sorted(input_map.items(), key=lambda item:(item[1]["hand_type_value"], item[1]["hand_value_by_character_priority"]), reverse=True))

    total_winnings = 0

    for key, value in sorted_by_character_priority.items():
        total_winnings += value["bid"] * hand_count
        hand_count -= 1

    print(total_winnings)