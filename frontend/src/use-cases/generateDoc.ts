export async function generateDoc(mrLink: string, gitlabToken: string) {
  const response = await fetch(`/generate-doc`, {
    method: "POST",
    body: JSON.stringify({ 
      mrLink: mrLink,
      gitlabToken: gitlabToken,
     }),
  });

  return response.json();
}
