import os
import json
# pip install Pillow
from PIL import Image

jsons_path = "./assets/char_jsons"
origs_path = "./assets/original"

diam_frame = "./assets/misc/frame_diam.png"
gold_frame = "./assets/misc/frame_gold.png"
silv_frame = "./assets/misc/frame_silv.png"

result_path = "./assets/framed"

width = 384
height = 640
border = 7

def frame_image(frame_path, orig_path):
    global width, height, border

    # open frame
    frame = Image.open(frame_path)
    frame = frame.convert("RGBA")

    # open original card
    orig = Image.open(orig_path)
    orig = orig.resize((width-border*2, height-border*2))
    orig = orig.convert("RGBA")

    # combine em
    framed_image = Image.new("RGBA", (width, height), (0, 0, 0, 0))
    framed_image.paste(orig, (border, border))
    framed_image = Image.alpha_composite(framed_image, frame)

    return framed_image

if __name__=="__main__":
    for filename in os.listdir(jsons_path):
        # load character's json file and determine the paths
        with open(f"{jsons_path}/{filename}") as fh:
            json_file = json.load(fh)

        if json_file["rarity"] == 5:
            frame_path = diam_frame
        if json_file["rarity"] == 4:
            frame_path = gold_frame
        if json_file["rarity"] == 3:
            frame_path = silv_frame

        orig_path = f"{origs_path}/{json_file['name']}.png"
        save_path = f"{result_path}/{json_file['name']}_framed.png"

        # frame the card and save it
        framed_card = frame_image(frame_path, orig_path)
        framed_card.save(save_path, "png")
