"use client";

import React from "react";
import TitleCard from "@/components/titleCard";
import Input from "@/components/ui/input";
import Card from "@/components/ui/card";
import Button from "@/components/ui/button";
import { useState } from "react";
import { generateDoc, generateDocFromAuthor } from "@/use-cases/generateDoc";
import { useDocContext } from "@/context/doc-context";
import Select from "@/components/ui/select";
import SelectItem from "@/components/ui/select-item";
import MdViewer from "@/components/mdViewer";
import { gitlabColors } from "@/lib/colors";
import ListInput from "@/components/ui/list-input";
import GenerateButton from "@/components/generate-button";

enum RequestType {
	MANUAL = "manual",
	AUTHOR = "author",
}

export default function Home() {
	const [mrLinks, setMrLinks] = useState([""]);
	const [model, setModel] = useState("gpt-4o");
	const [modelInput, setModelInput] = useState("");
	const [modelInputOpen, setModelInputOpen] = useState(false);
	const [requestType, setRequestType] = useState(RequestType.MANUAL);
	const [gitlabUsername, setGitlabUsername] = useState("");
	const [status, setStatus] = useState("");
	const [errors, setErrors] = useState({
		mrLink: "",
		gitlabUsername: "",
	});
	const [loading, setLoading] = useState(false);
	const { doc, setDoc } = useDocContext();
	const [copySuccess, setCopySuccess] = useState(false);

	const validateForm = () => {
		let isValid = true;
		const newErrors = { mrLink: "", gitlabUsername: "" };

		if (requestType === RequestType.MANUAL) {
			if (mrLinks.length === 0) {
				setStatus("MR link is required");
				isValid = false;
			} else if (!mrLinks.every((link: string) => link.includes("gitlab"))) {
				setStatus("At least one MR link is invalid, only GitLab MR links are supported as of now.");
				isValid = false;
			}
		} else {
			if (!gitlabUsername.trim()) {
				setStatus("GitLab username is required");
				isValid = false;
			}
		}

		setErrors(newErrors);
		return isValid;
	};

	const handleGenerate = async () => {
		if (validateForm()) {
			try {
				setLoading(true);
				if (requestType === RequestType.MANUAL) {
					const response: string = await generateDoc(mrLinks, model, setStatus);
					console.log(response);
					setDoc(response);
				} else {
					const response: string = await generateDocFromAuthor(gitlabUsername, model, setStatus);
					console.log(response);
					setDoc(response);
				}
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
			<Card className="w-full max-w-xs sm:max-w-md md:max-w-lg lg:max-w-xl gap-0">
				<TitleCard className="mb-10" companyName="GitLab" companyColors={Object.values(gitlabColors)} />
				<div className="w-full flex flex-col sm:flex-row gap-2">
					<div className="flex-[4] flex flex-col gap-2">

						<div className="flex flex-row">
							<Button className={`w-full py-1 ${requestType === RequestType.MANUAL ? "bg-[#e64a19]/80" : "bg-transparent"}`} onClick={() => setRequestType(RequestType.MANUAL)}>MR Link(s)</Button>
							<Button className={`w-full py-1 ${requestType === RequestType.AUTHOR ? "bg-[#e64a19]/80" : "bg-transparent"}`} onClick={() => setRequestType(RequestType.AUTHOR)}>Author</Button>
						</div>
						
						{requestType === RequestType.MANUAL ? (
							<>
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
							</>
						) : (
							<>
								<Input
									label="GitLab Username"
									value={gitlabUsername}
									onChange={(e) => setGitlabUsername(e.target.value)}
									placeholder="Enter your GitLab username"
									className={errors.gitlabUsername ? "border-red-500" : ""} 
								/>
								{errors.gitlabUsername && (
									<span className="text-red-500 text-sm">{errors.gitlabUsername}</span>
								)}
							</>
						)}

						<div className="flex flex-row gap-2">
							{modelInputOpen ? (
								<>
									<Input
										label="Ollama model"
										value={modelInput}
										className="w-full"
										onChange={(e) => setModelInput(e.target.value)}
										placeholder="Enter an Ollama model"
									/>
									<Button className="text-xs px-1 h-10" onClick={() => setModelInputOpen(!modelInputOpen)}>use default models</Button>
								</>
							): (
								<>
									<Select
										value={model}
										onChange={(e) => setModel(e.target.value)}
									>
										<SelectItem value="gpt-4o">GPT-4o</SelectItem>
										<SelectItem value="llama3.2">Llama 3.2 (Ollama)</SelectItem>
									</Select>
									<Button className="text-xs px-1 h-10" onClick={() => setModelInputOpen(!modelInputOpen)}>input a model</Button>
								</>
							)}
						</div>
					</div>

					<GenerateButton 
						loading={loading}
						onClick={handleGenerate}
					/>

				</div>
				{status && (
					<p className="mt-2 text-red-500 text-xs">{status}</p>
				)}
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

