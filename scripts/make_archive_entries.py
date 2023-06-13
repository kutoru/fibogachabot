import os
import json
# pip install Pillow
from PIL import Image, ImageFont, ImageDraw

jsons_path = "./assets/char_jsons"
framed_path = "./assets/framed"

diam_bg = "./assets/misc/arch_entry_diam.png"
gold_bg = "./assets/misc/arch_entry_gold.png"
silv_bg = "./assets/misc/arch_entry_silv.png"

diam_stars = "./assets/misc/stars_diam.png"
gold_stars = "./assets/misc/stars_gold.png"
silv_stars = "./assets/misc/stars_silv.png"

card_shadow = "./assets/misc/arch_card_shadow.png"

font_path = "./assets/misc/myriad_pro.otf"
font_color = (255, 255, 255, 255)

result_path = "./assets/arch_entries"

def make_archive_entry(bg_path, card_path, star_path, name_str, nickname_str, description_str):
    global card_shadow

    # open bg image and get its width and height
    arch_entry = Image.open(bg_path).convert("RGBA")
    width, height = arch_entry.size

    # add the card to the image
    card = Image.open(card_path)
    arch_entry.paste(card, (30, 30))

    # add the shadow to the image
    shadow = Image.open(card_shadow)
    positioned_shadow = Image.new("RGBA", (width, height), (0, 0, 0, 0))
    positioned_shadow.paste(shadow, (30, 30))
    arch_entry = Image.alpha_composite(arch_entry, positioned_shadow)

    # add the stars
    stars = Image.open(star_path)
    positioned_stars = Image.new("RGBA", (width, height), (0, 0, 0, 0))
    positioned_stars.paste(stars, (121, 617))
    arch_entry = Image.alpha_composite(arch_entry, positioned_stars)

    # add the name
    add_text_to_image(arch_entry, name_str, (451+15-3, 27+15-2), 60, 392)  # actual font size = font size * 1.125

    # add the nickname
    add_text_to_image(arch_entry, nickname_str, (451+15, 27+15+54+15-1), 30, 392)

    # add the description
    add_text_to_image(arch_entry, description_str, (451+15, 187+15-1), 25, 392)

    return arch_entry

def add_text_to_image(image, text, box, font_size, wrap_length):
    global font_path, font_color

    # get the font
    font = ImageFont.truetype(font_path, font_size)

    # wrap the text
    text = get_wrapped_text(text, font, wrap_length)

    # draw the text on the image
    image_draw = ImageDraw.Draw(image)
    image_draw.text(box, text, font=font, fill=font_color)

def get_wrapped_text(text, font, max_line_length):  # max length in pixels
    lines = [""]

    for word in text.split():
        new_line = f"{lines[-1]} {word}".strip()

        if font.getlength(new_line) > max_line_length:
            lines.append(word)
        else:
            lines[-1] = new_line

    return "\n".join(lines)

if __name__=="__main__":
    for filename in os.listdir(jsons_path):
        # load character's json file and determine the necessary info
        with open(f"{jsons_path}/{filename}") as fh:
            json_file = json.load(fh)

        name = json_file["name"]
        nickname = json_file["nickname"]
        description = json_file["description"]
        rarity = json_file["rarity"]

        if len(nickname) == 0:
            nickname = "No nickname"

        if len(description) == 0:
            description = "No description"

        if rarity == 5:
            bg_path = diam_bg
            star_path = diam_stars
        if rarity == 4:
            bg_path = gold_bg
            star_path = gold_stars
        if rarity == 3:
            bg_path = silv_bg
            star_path = silv_stars

        card_path = f"{framed_path}/{name}_framed.png"
        save_path = f"{result_path}/{name}_arch_entry.png"

        # make the image and save it
        archive_entry = make_archive_entry(bg_path, card_path, star_path, name, nickname, description)
        archive_entry.save(save_path, "png")
