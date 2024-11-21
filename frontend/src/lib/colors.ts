export const gitlabColors = {
  BUCKTHORN: "#fca326",
  ORANGE: "#fc6d26",
  CINNABAR: "#e24329",
};

export function splitTextIntoColorChunks(text: string, colors: string[]): { text: string, color: string }[] {
  const chunkSize = Math.ceil(text.length / colors.length);
  const chunks = text.match(new RegExp(`.{1,${chunkSize}}`, 'g')) || [];
  
  return chunks.map((chunk, index) => ({
    text: chunk,
    color: colors[index] || colors[colors.length - 1] // fallback to last color if needed
  }));
}