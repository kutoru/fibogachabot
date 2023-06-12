from PIL import Image, ImageFont
from pathlib import Path
from os import listdir

# paths
CURRENT_DIR = str(Path(__file__).parent.resolve())

bg_path = f"{CURRENT_DIR}\\Assets\\arch_prev_bg.png"

diam_frame = f"{CURRENT_DIR}\\Assets\\pfp_frame_diam.png"
gold_frame = f"{CURRENT_DIR}\\Assets\\pfp_frame_gold.png"
silv_frame = f"{CURRENT_DIR}\\Assets\\pfp_frame_silv.png"

diam_pfps = f"{CURRENT_DIR}\\5 star"
gold_pfps = f"{CURRENT_DIR}\\4 star"
silv_pfps = f"{CURRENT_DIR}\\3 star"

# text
uppercase = False
cut_descenders = True
font_stroke = 0
text_color = (255, 255, 255, 255)

font1 = [f"{CURRENT_DIR}\\Assets\\Myriad Pro.otf", 20, 2]  # too long
font2 = [f"{CURRENT_DIR}\\Assets\\Permanent Headline.otf", 24, 2]  # too thin
font3 = [f"{CURRENT_DIR}\\Assets\\Poppins.otf", 24, 2]  # too long
font4 = [f"{CURRENT_DIR}\\Assets\\Crushed.ttf", 24, 2]  # shit
font5 = [f"{CURRENT_DIR}\\Assets\\False Positive.ttf", 24, 2]  # shit
font6 = [f"{CURRENT_DIR}\\Assets\\Wayland.otf", 22, 1]  # current font

for index, item in enumerate(font6):
    if index == 0: font_file_path = item
    elif index == 1: font_size = item
    elif index == 2: font_y_offset = item

# cell
number_of_repeats = 0
x_offset = 3
y_offset = 3
pfp_width = 125
pfp_height = 125
name_height = 20
border_size = 8
cell_width = pfp_width + border_size * 2
cell_height = pfp_height + name_height + border_size * 2

# canvas
x_cells = 9
y_cells = 6
canvas_width = cell_width * x_cells + x_offset * 2
canvas_height = cell_height * y_cells + y_offset * 2

def make_cell(border_path, pfp_path):
    global cut_descenders, uppercase
    border = Image.open(border_path).convert("RGBA")

    pfp = Image.open(pfp_path).convert("RGBA")
    pfp = pfp.resize((pfp_width, pfp_height))

    name_string = pfp_path.split("\\")[-1].split("_")[0]
    if uppercase: name_string = name_string.upper()
    name = make_name(name_string)

    border.paste(pfp, (border_size, border_size))
    new_name = Image.new("RGBA", (cell_width, cell_height), (0, 0, 0, 0))
    new_name.paste(name, (border_size, border_size+pfp_height))
    border = Image.alpha_composite(border, new_name)

    return border

def make_name(name_string):
    global font_size, font_stroke, font_file_path, text_color, font_y_offset, cut_descenders

    font = ImageFont.truetype(font_file_path, size=font_size)
    mask_image = font.getmask(name_string, "L", stroke_width=font_stroke)

    text_image = Image.new("RGBA", mask_image.size)
    text_image.im.paste(text_color, (0, 0) + mask_image.size, mask_image)

    mask_x, mask_y = mask_image.size
    x_offset = (125 - mask_x) // 2
    y_offset = font_y_offset

    name = Image.new("RGBA", (125, 20), (0, 0, 0, 0))
    #name = Image.new("RGBA", (125, 25), (0, 0, 0, 0))

    name.paste(text_image, (x_offset, y_offset))
    return name

def fill_canvas(canvas):
    global number_of_repeats
    char_list = []

    for index, dir in enumerate((diam_pfps, gold_pfps, silv_pfps)):
        for i in range(number_of_repeats+1):  # test
            for file_name in listdir(dir):
                if file_name.endswith("_pfp.png"):
                    char_dict = {"pfp_path": f"{dir}\\{file_name}"}

                    if index == 0:
                        char_dict["border_path"] = diam_frame
                    elif index == 1:
                        char_dict["border_path"] = gold_frame
                    elif index == 2:
                        char_dict["border_path"] = silv_frame

                    char_list.append(char_dict)

    for y in range(y_cells):
        for x in range(x_cells):
            index = y * x_cells + x
            if (index+1) > len(char_list) or (index+1) > (y_cells*x_cells):
                break

            cell = make_cell(char_list[index]["border_path"], char_list[index]["pfp_path"])

            total_x = x_offset + cell_width * x
            total_y = y_offset + cell_height * y
            cell_canvas = Image.new("RGBA", (canvas_width, canvas_height), (0, 0, 0, 0))
            cell_canvas.paste(cell, (total_x, total_y))

            canvas = Image.alpha_composite(canvas, cell_canvas)

    return canvas

def main():
    canvas = Image.open(bg_path).convert("RGBA")
    canvas = fill_canvas(canvas)
    #canvas.show()
    canvas.save(f"{CURRENT_DIR}\\Archive preview.png", "png")

main()
