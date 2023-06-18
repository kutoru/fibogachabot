import os
import zipfile

original_folders = ["char_jsons", "faces", "misc", "original", "other", "gift_images", "quest_jsons"]
original_files = ["gifts.json"]

generatable_folders = ["arch_entries", "framed"]
generatable_files = ["archive_preview.png"]

output_path = "./assets/assets.zip"

def archive_all_assets():
    zh = zipfile.ZipFile(output_path, "w")
    zh.write("assets")

    for name in original_folders + original_files + generatable_folders + generatable_files:
        zh.write(f"assets/{name}")

    for foldername in original_folders + generatable_folders:
        for filename in os.listdir(f"./assets/{foldername}"):
            zh.write(f"assets/{foldername}/{filename}")

    zh.close()

def archive_original_assets():
    zh = zipfile.ZipFile(output_path, "w")
    zh.write("assets")

    for name in original_folders + original_files:
        zh.write(f"assets/{name}")

    for foldername in original_folders:
        for filename in os.listdir(f"./assets/{foldername}"):
            zh.write(f"assets/{foldername}/{filename}")

    zh.close()

if __name__=="__main__":
    user_input = input("Archive generatable files? (Generatable files are files that can be created with scripts) (y/n):\n").strip()

    if user_input == "y":
        archive_all_assets()
    elif user_input == "n":
        archive_original_assets()
    else:
        print("Invalid input")
