const BACKEND_URL = process.env.NEXT_PUBLIC_BACKEND_URL || "http://localhost:8080"

if (BACKEND_URL === "http://localhost:8080") {
	console.log("WARNING: using default backend URL, this is for development only")
}

export { BACKEND_URL }
