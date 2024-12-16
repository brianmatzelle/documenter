export async function generateDoc(mrLinks: string[], gitlabToken: string, model: string, setStatus: (status: string) => void): Promise<string> {
  // Return an EventSource that the caller can use to listen for events
  const eventSource = new EventSource(`/generate-doc?${new URLSearchParams({
    mrLinks: JSON.stringify(mrLinks),
    gitlabToken,
    model: model.toLowerCase(),
  })}`);

  return new Promise((resolve, reject) => {
    eventSource.onmessage = (event) => {
      const data = JSON.parse(event.data);
      setStatus(data.message);
    };

    eventSource.addEventListener('status', (event) => {
      console.log('Status update:', event.data);
      const data = JSON.parse(event.data);
      setStatus(data.message);
    });

    eventSource.addEventListener('complete', (event) => {
      const data = JSON.parse(event.data);
      eventSource.close();
      setStatus("");
      resolve(data.doc);
    });

    eventSource.addEventListener('error', (event: MessageEvent) => {
      const data = JSON.parse(event.data);
      eventSource.close();
      setStatus(data.error);
      reject(new Error(data.error));
    });
  });
}

export async function generateDocFromAuthor(gitlabUsername: string, gitlabToken: string, model: string, setStatus: (status: string) => void): Promise<string> {
  const eventSource = new EventSource(`/gen-from-author?${new URLSearchParams({
    author: gitlabUsername,
    gitlabToken,
    model: model.toLowerCase(),
  })}`);

  return new Promise((resolve, reject) => {
    eventSource.onmessage = (event) => {
      const data = JSON.parse(event.data);
      setStatus(data.message);
    };

    eventSource.addEventListener('status', (event) => {
      const data = JSON.parse(event.data);
      setStatus(data.message);
    });

    eventSource.addEventListener('complete', (event) => {
      const data = JSON.parse(event.data);
      eventSource.close();
      setStatus("");
      resolve(data.doc);
    });

    eventSource.addEventListener('error', (event: MessageEvent) => {
      const data = JSON.parse(event.data);
      eventSource.close();
      setStatus(data.error);
      reject(new Error(data.error));
    });
  });
}
