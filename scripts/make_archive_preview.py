import os
import json
# pip install Pillow
from PIL import Image, ImageFont

# paths
pfps_path = "./assets/faces"
jsons_path = "./assets/char_jsons"
bg_path = "./assets/misc/arch_prev_bg.png"

diam_frame = "./assets/misc/pfp_frame_diam.png"
gold_frame = "./assets/misc/pfp_frame_gold.png"
silv_frame = "./assets/misc/pfp_frame_silv.png"

result_path = "./assets/archive_preview.png"

# text
text_color = (255, 255, 255, 255)

font_file_path = "./assets/misc/wayland.otf"
font_size = 22
font_y_offset = 1

# cell
border_size = 8
pfp_width = 125
pfp_height = 125
name_height = 20
name_width = pfp_width

cell_width = pfp_width + (border_size * 2)
cell_height = pfp_height + name_height + (border_size * 2)

# canvas
x_offset = 3
y_offset = 3
x_cells = 9
y_cells = 6

canvas_width = cell_width * x_cells + x_offset * 2
canvas_height = cell_height * y_cells + y_offset * 2

def make_cell(border_path, pfp_path, name_string):
    global pfp_width, pfp_height, cell_width, cell_height, border_size

    # load border
    cell = Image.open(border_path).convert("RGBA")

    # load and resize the face
    pfp = Image.open(pfp_path).convert("RGBA")
    pfp = pfp.resize((pfp_width, pfp_height))

    # convert name string to an image and position it correctly
    temp_name = make_name(name_string)
    name = Image.new("RGBA", (cell_width, cell_height), (0, 0, 0, 0))
    name.paste(temp_name, (border_size, border_size+pfp_height))

    # combine 'em all
    cell.paste(pfp, (border_size, border_size))
    cell = Image.alpha_composite(cell, name)

    return cell

def make_name(name_string):
    global font_file_path, font_size, text_color, font_y_offset

    # make the mask
    font = ImageFont.truetype(font_file_path, size=font_size)
    mask_image = font.getmask(name_string, "L", stroke_width=0)

    # create an image and paste the mask
    text_image = Image.new("RGBA", mask_image.size)
    text_image.im.paste(text_color, (0, 0) + mask_image.size, mask_image)

    # position the text correctly
    mask_x, mask_y = mask_image.size
    text_x_offset = (125 - mask_x) // 2
    text_y_offset = font_y_offset

    name = Image.new("RGBA", (125, 20), (0, 0, 0, 0))
    name.paste(text_image, (text_x_offset, text_y_offset))

    return name

def get_chars_info():
    global jsons_path, diam_frame, gold_frame, silv_frame, pfps_path

    chars_info = []

    # load the necessary character info into chars_info for all characters
    for filename in os.listdir(jsons_path):
        with open(f"{jsons_path}/{filename}") as fh:
            json_file = json.load(fh)

        if json_file["rarity"] == 5:
            border_path = diam_frame
        elif json_file["rarity"] == 4:
            border_path = gold_frame
        elif json_file["rarity"] == 3:
            border_path = silv_frame

        pfp_path = f"{pfps_path}/{json_file['name']}_pfp.png"

        chars_info.append([border_path, pfp_path, json_file['name'], json_file['id']])

    return chars_info

def fill_canvas(canvas):
    global x_cells, y_cells, x_offset, y_offset, cell_width, cell_height, canvas_width, canvas_height

    chars_info = get_chars_info()

    for index, char_info in enumerate(chars_info):
        if index == (x_cells * y_cells):
            break

        # create the archive cell
        cell = make_cell(char_info[0], char_info[1], char_info[2])

        # calculate the cell position
        curr_row = index // x_cells
        curr_col = index - (curr_row * x_cells)
        total_x = x_offset + cell_width * curr_col
        total_y = y_offset + cell_height * curr_row

        # position the cell correctly
        cell_canvas = Image.new("RGBA", (canvas_width, canvas_height), (0, 0, 0, 0))
        cell_canvas.paste(cell, (total_x, total_y))

        # add the cell to the canvas
        canvas = Image.alpha_composite(canvas, cell_canvas)

    return canvas

if __name__=="__main__":
    canvas = Image.open(bg_path).convert("RGBA")
    canvas = fill_canvas(canvas)
    canvas.save(result_path, "png")
