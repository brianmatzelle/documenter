from PIL import Image, ImageSequence
import os

INPUT_GIF = "documenter-demo.gif"
OUTPUT_GIF = "documenter-demo-compressed.gif"

def compress_gif(input_path, output_path, max_colors=128, optimize=True):
    """
    Compress a GIF by reducing the number of colors and optimizing.

    Parameters:
        input_path (str): Path to the input GIF file.
        output_path (str): Path to save the compressed GIF file.
        max_colors (int): Maximum number of colors to reduce to (default is 128).
        optimize (bool): Whether to optimize the GIF (default is True).
    """
    try:
        # Open the input GIF
        with Image.open(input_path) as gif:
            frames = []
            total_frames = gif.n_frames  # Get total number of frames
            print(f"Processing {total_frames} frames...")

            for i, frame in enumerate(ImageSequence.Iterator(gif)):
                print(f"Converting frame {i+1}/{total_frames}")
                frame = frame.convert("P", palette=Image.ADAPTIVE, colors=max_colors)
                frames.append(frame)

            print("Saving compressed GIF...")
            # Save the compressed GIF
            frames[0].save(
                output_path,
                save_all=True,
                append_images=frames[1:],
                optimize=optimize,
                loop=gif.info.get("loop", 0),
                duration=gif.info.get("duration", 100),
                disposal=gif.info.get("disposal", 2)
            )

        print(f"Compressed GIF saved at {output_path}")

    except Exception as e:
        print(f"Error compressing GIF: {e}")

if __name__ == "__main__":
    # Example usage
    input_gif = INPUT_GIF
    output_gif = OUTPUT_GIF

    # Ensure the input file exists
    if not os.path.exists(input_gif):
        print(f"Error: The file {input_gif} does not exist.")
    else:
        compress_gif(input_gif, output_gif)