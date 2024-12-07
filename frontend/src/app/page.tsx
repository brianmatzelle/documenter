"use client";

import TitleCard from "@/components/titleCard";
import Input from "@/components/ui/input";
import Card from "@/components/ui/card";
import Button from "@/components/ui/button";
import { useState } from "react";
import { generateDoc } from "@/use-cases/generateDoc";
import { useDocContext } from "@/context/doc-context";
import Select from "@/components/ui/select";
import SelectItem from "@/components/ui/select-item";
import MdViewer from "@/components/mdViewer";
import { gitlabColors } from "@/lib/colors";
import ListInput from "@/components/ui/list-input";
import GenerateButton from "@/components/generate-button";

export default function Home() {
	// const [mrLink, setMrLink] = useState("");
	const [mrLinks, setMrLinks] = useState([""]);
	const [gitlabToken, setGitlabToken] = useState("");
	const [model, setModel] = useState("openai");
	const [errors, setErrors] = useState({
		mrLink: "",
		gitlabToken: "",
	});
	const [loading, setLoading] = useState(false);
	const { doc, setDoc } = useDocContext();
	const [copySuccess, setCopySuccess] = useState(false);

	const validateForm = () => {
		let isValid = true;
		const newErrors = { mrLink: "", gitlabToken: "" };

		if (mrLinks.length === 0) {
			newErrors.mrLink = "MR link is required";
			isValid = false;
		} else if (!mrLinks.every(link => link.includes("gitlab"))) {
			newErrors.mrLink = "At least one MR link is invalid, only GitLab MR links are supported as of now.";
			isValid = false;
		}

		if (!gitlabToken.trim()) {
			newErrors.gitlabToken = "GitLab token is required";
			isValid = false;
		}

		setErrors(newErrors);
		return isValid;
	};

	const handleGenerate = async () => {
		if (validateForm()) {
			try {
				setLoading(true);
				const response = await generateDoc(mrLinks, gitlabToken, model);
				console.log(response);
				setDoc(response.doc);
			} catch (error) {
				console.error("Error generating doc:", error);
			} finally {
				setLoading(false);
			}
		}
	};

	const handleCopyDoc = async () => {
		try {
			await navigator.clipboard.writeText(doc);
			setCopySuccess(true);
			setTimeout(() => setCopySuccess(false), 2000);
		} catch (err) {
			console.error('Failed to copy text:', err);
		}
	};

	return (
		<div className="flex-1 flex flex-col items-center justify-around gap-4 p-4">
			<Card className="w-full max-w-xs sm:max-w-md md:max-w-lg lg:max-w-xl">
				<TitleCard companyName="GitLab" companyColors={Object.values(gitlabColors)} />
				<div className="w-full flex flex-col sm:flex-row gap-2">
					<div className="flex-[4] flex flex-col gap-2">
						
						<ListInput 
							mrLinks={mrLinks}
							onMrLinksChange={setMrLinks}
							label="Merge Request Link"
							placeholder="Enter your MR link here"
							className={errors.mrLink ? "border-red-500" : ""} 
						/>
						{errors.mrLink && (
							<span className="text-red-500 text-sm">{errors.mrLink}</span>
						)}

						<Input 
							label="GitLab Token"
							value={gitlabToken}
							onChange={(e) => {
								setGitlabToken(e.target.value);
								if (errors.gitlabToken) setErrors(prev => ({ ...prev, gitlabToken: "" }));
							}}
							placeholder="Enter your GitLab token here"
							className={errors.gitlabToken ? "border-red-500" : ""} 
							type="password"
						/>
						{errors.gitlabToken && (
							<span className="text-red-500 text-sm">{errors.gitlabToken}</span>
						)}

						<Select
							value={model}
							onChange={(e) => setModel(e.target.value)}
						>
							<SelectItem value="openai">GPT-4o</SelectItem>
							<SelectItem value="ollama">Llama 3.2 (Ollama)</SelectItem>
						</Select>
					</div>

					<GenerateButton 
						loading={loading}
						onClick={handleGenerate}
					/>

				</div>
			</Card>

			{doc && (
				<div className="w-full max-w-xs sm:max-w-md md:max-w-lg lg:max-w-xl">
					<Button 
						onClick={handleCopyDoc}
						className="mb-2 w-full"
					>
						{copySuccess ? "Copied!" : "Copy Document"}
					</Button>
					<MdViewer content={doc} />
				</div>
			) || <div>No document generated yet</div>}
		</div>
	)
}

