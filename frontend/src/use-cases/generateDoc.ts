export async function generateDoc(mrLink: string, gitlabToken: string, model: string) {
  const response = await fetch(`/generate-doc`, {
    method: "POST",
    body: JSON.stringify({ 
      mrLink: mrLink,
      gitlabToken: gitlabToken,
      model: model,
     }),
  });

  return response.json();
}
