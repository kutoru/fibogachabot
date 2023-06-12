from PIL import Image
from pathlib import Path
from os import listdir

CURRENT_DIR = str(Path(__file__).parent.resolve())

diam_path = f"{CURRENT_DIR}\\5 star"
gold_path = f"{CURRENT_DIR}\\4 star"
silv_path = f"{CURRENT_DIR}\\3 star"

diam_frame = f"{CURRENT_DIR}\\Assets\\frame_diam.png"
gold_frame = f"{CURRENT_DIR}\\Assets\\frame_gold.png"
silv_frame = f"{CURRENT_DIR}\\Assets\\frame_silv.png"

width = 384
height = 640
border = 7

def combine_images(image1_path, image2_path, save_path):
    image1 = Image.open(image1_path)
    image1 = image1.resize((width-border*2, height-border*2))
    image1 = image1.convert("RGBA")

    image2 = Image.open(image2_path)
    image2 = image2.convert("RGBA")

    new_image = Image.new("RGBA", (width, height), (0, 0, 0, 0))
    new_image.paste(image1, (border, border))
    new_image = Image.alpha_composite(new_image, image2)
    new_image.save(save_path, "png")

def main():
    for index, dir in enumerate((diam_path, gold_path, silv_path)):
        for file_name in listdir(dir):
            if "_" not in file_name:
                if index == 0:
                    frame_path = diam_frame
                    new_image_path = diam_path
                elif index == 1:
                    frame_path = gold_frame
                    new_image_path = gold_path
                elif index == 2:
                    frame_path = silv_frame
                    new_image_path = silv_path

                new_image_path += "\\" + file_name.split(".")[0] + "_framed.png"

                combine_images(f"{dir}\\{file_name}", frame_path, new_image_path)

main()
