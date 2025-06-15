import re

with open("meaning_of_life.txt", "r") as f:
    content = f.read()

# remove all the text between [ and ] there could be multiple lines between the []
content = re.sub(r"\[.+?\]", "", content, flags=re.DOTALL)

# remove all the text between ( and )
content = re.sub(r"\(.+?\)", "", content, flags=re.DOTALL)

content = re.sub(r" +", " ", content)
content = re.sub(r"-+", "-", content)

content = (
    content.lower()
    .replace("...", "\n")
    .replace(".", "\n")
    .replace(",", "\n")
    .replace("!", "\n")
    .replace("?", "\n")
    .replace(";", "\n")
    .replace(":", "\n")
    .replace(" - ", "\n")
    # .replace("'", "\n")
    .replace("*", "")
    .replace("_", "")
    .split("\n")
)

content = [x.strip() for x in content]

content = list(filter(lambda x: x not in ["", "in", "and"], content))

# make a list of unique elements
s = set()
new_content = []
for x in content:
    if x not in s:
        s.add(x)
        new_content.append(x)

# write the list to a file
with open("meaning_of_life_clean.txt", "w") as f:
    f.write("\n".join(new_content))
