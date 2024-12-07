export async function generateDoc(mrLinks: string[], gitlabToken: string, model: string) {
  const response = await fetch(`/generate-doc`, {
    method: "POST",
    body: JSON.stringify({ 
      mrLinks: mrLinks,
      gitlabToken: gitlabToken,
      model: model,
     }),
  });

  return response.json();
}
