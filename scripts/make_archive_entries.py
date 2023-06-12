from PIL import Image, ImageFont, ImageDraw
from pathlib import Path
from os import listdir
import sqlite3

CURRENT_DIR = str(Path(__file__).parent.resolve())

diam_path = f"{CURRENT_DIR}\\5 star"
gold_path = f"{CURRENT_DIR}\\4 star"
silv_path = f"{CURRENT_DIR}\\3 star"

diam_bg = f"{CURRENT_DIR}\\Assets\\arch_entry_diam.png"
gold_bg = f"{CURRENT_DIR}\\Assets\\arch_entry_gold.png"
silv_bg = f"{CURRENT_DIR}\\Assets\\arch_entry_silv.png"

diam_stars = f"{CURRENT_DIR}\\Assets\\stars_diam.png"
gold_stars = f"{CURRENT_DIR}\\Assets\\stars_gold.png"
silv_stars = f"{CURRENT_DIR}\\Assets\\stars_silv.png"

card_shadow = f"{CURRENT_DIR}\\Assets\\arch_card_shadow.png"

font_path = f"{CURRENT_DIR}\\Assets\\Myriad Pro.otf"

def combine_elements(bg_path, card_path, star_path, name_str, nickname_str, description_str):
    # adding bg
    arch_entry = Image.open(bg_path).convert("RGBA")
    width, height = arch_entry.size

    # adding card
    card = Image.open(card_path)
    arch_entry.paste(card, (30, 30))

    # adding card shadow
    shadow = Image.open(card_shadow)
    extended_shadow = Image.new("RGBA", (width, height), (0, 0, 0, 0))
    extended_shadow.paste(shadow, (30, 30))
    arch_entry = Image.alpha_composite(arch_entry, extended_shadow)

    # adding stars
    stars = Image.open(star_path)
    extended_stars = Image.new("RGBA", (width, height), (0, 0, 0, 0))
    extended_stars.paste(stars, (121, 617))
    arch_entry = Image.alpha_composite(arch_entry, extended_stars)

    # adding name
    add_text_to_image(arch_entry, name_str, (451+15-3, 27+15-2), 60, 392)  # actual font size = font size * 1.125

    # adding nickname
    add_text_to_image(arch_entry, nickname_str, (451+15, 27+15+54+15-1), 30, 392)

    # adding description
    add_text_to_image(arch_entry, description_str, (451+15, 187+15-1), 25, 392)

    return arch_entry

def add_text_to_image(image_obj, text, box, font_size, wrap_length):
    font = ImageFont.truetype(font_path, size=font_size)
    text = get_wrapped_text(text, font, wrap_length)
    image_draw = ImageDraw.Draw(image_obj)
    image_draw.text(box, text, font=font, fill=(255, 255, 255, 255))
    #return image_obj

def get_wrapped_text(text, font, max_line_length):  # max length in pixels
    lines = [""]
    for word in text.split():
        new_line = f"{lines[-1]} {word}".strip()
        if font.getlength(new_line) > max_line_length:
            lines.append(word)
        else:
            lines[-1] = new_line
    return "\n".join(lines)

def main():
    for index, dir in enumerate((diam_path, gold_path, silv_path)):
        for file_name in listdir(dir):
            if file_name.endswith("_framed.png"):
                if index == 0:
                    bg_path = diam_bg
                    card_path = diam_path + "\\" + file_name
                    star_path = diam_stars
                    new_image_path = diam_path
                elif index == 1:
                    bg_path = gold_bg
                    card_path = gold_path + "\\" + file_name
                    star_path = gold_stars
                    new_image_path = gold_path
                elif index == 2:
                    bg_path = silv_bg
                    card_path = silv_path + "\\" + file_name
                    star_path = silv_stars
                    new_image_path = silv_path

                name_str = file_name.split("_")[0]
                con = sqlite3.connect("database.db")
                cur = con.cursor()
                cur.execute("select nickname, description from characters where name = ?;", (name_str,))
                nickname_str, description_str = cur.fetchone()

                arch_entry = combine_elements(bg_path, card_path, star_path, name_str, nickname_str, description_str)
                new_image_path += "\\" + file_name.split("_")[0] + "_arch_entry.png"
                arch_entry.save(new_image_path, "png")

main()
