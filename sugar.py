#!/usr/bin/env python

import os
import sys
import argparse
from dotenv import load_dotenv
from openai import OpenAI

# Load environment variables from .env file
# https://github.com/pyinstaller/pyinstaller/issues/5522#issuecomment-770858489
extDataDir = os.getcwd()
if getattr(sys, 'frozen', False):
    extDataDir = sys._MEIPASS
load_dotenv(dotenv_path=os.path.join(extDataDir, '.env'))

# Get the OpenAI API key from environment variables
client = OpenAI(
    # This is the default and can be omitted
    api_key=os.environ.get("OPENAI_API_KEY"),
)

def make_sentence_more_friendly(sentence):
    prompt = "Make the user-provided sentence less blunt, and a bit more friendly, but without exaggeration"
    response = client.chat.completions.create(
        messages=[
            {
                "role": "system",
                "content": prompt
            },
            {
                "role": "user",
                "content": sentence,
            }
        ],
        model="gpt-3.5-turbo",
    )
    return response.choices[0].message.content.strip()

def main():
    parser = argparse.ArgumentParser(description="Make a blunt sentence a bit more friendly")
    parser.add_argument("sentence", help="The blunt sentence to be made more friendly")
    args = parser.parse_args()

    friendly_sentence = make_sentence_more_friendly(args.sentence)
    print(friendly_sentence)

if __name__ == "__main__":
    main()
