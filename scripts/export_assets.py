import os
import zipfile

original_folders = ["char_jsons", "faces", "misc", "original", "other"]
generatable_folders = ["arch_entries", "framed"]
output_path = "./assets/assets.zip"

def archive_all_assets():
    zh = zipfile.ZipFile(output_path, "w")
    zh.write("assets")

    for name in os.listdir("./assets"):
        if not os.path.isdir(f"./assets/{name}") and not name.endswith(".zip"):
            zh.write(f"assets/{name}")

    for foldername in original_folders + generatable_folders:
        for filename in os.listdir(f"./assets/{foldername}"):
            zh.write(f"assets/{foldername}/{filename}")

    zh.close()

def archive_original_assets():
    zh = zipfile.ZipFile(output_path, "w")
    zh.write("assets")

    for foldername in generatable_folders:
        zh.write(f"assets/{foldername}")

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
