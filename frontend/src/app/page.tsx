"use client";

import TitleCard from "@/components/titleCard";
import Input from "@/components/ui/input";
import Card from "@/components/ui/card";
import Button from "@/components/ui/button";
import { useState } from "react";
import { generateDoc } from "@/use-cases/generateDoc";
import { useDocContext } from "@/context/doc-context";
import { HashLoader } from "react-spinners";
import Select from "@/components/ui/select";
import SelectItem from "@/components/ui/select-item";

export default function Home() {
	const [mrLink, setMrLink] = useState("");
	const [gitlabToken, setGitlabToken] = useState("");
	const [model, setModel] = useState("openai");
	const [errors, setErrors] = useState({
		mrLink: "",
		gitlabToken: "",
	});
	const [loading, setLoading] = useState(false);
	const { setDoc } = useDocContext();


	const validateForm = () => {
		let isValid = true;
		const newErrors = { mrLink: "", gitlabToken: "" };

		if (!mrLink.trim()) {
			newErrors.mrLink = "MR link is required";
			isValid = false;
		} else if (!mrLink.includes("gitlab")) {
			newErrors.mrLink = "Please enter a valid GitLab MR link";
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
				const response = await generateDoc(mrLink, gitlabToken);
				console.log(response);
				setDoc(response.doc);
			} catch (error) {
				console.error("Error generating doc:", error);
			} finally {
				setLoading(false);
			}
		}
	};

	return (
		<div className="flex flex-col items-center justify-around h-full p-4">
			<Card className="w-full max-w-xs sm:max-w-md md:max-w-lg lg:max-w-xl">
				<TitleCard />
				<div className="w-full flex flex-col sm:flex-row gap-2">
					<div className="flex-[4] flex flex-col gap-2">
						<div className="flex flex-col gap-1">
							<Input 
								value={mrLink}
								onChange={(e) => {
									setMrLink(e.target.value);
									if (errors.mrLink) setErrors(prev => ({ ...prev, mrLink: "" }));
								}}
								placeholder="Enter your MR link here"
								className={errors.mrLink ? "border-red-500" : ""} 
							/>
							{errors.mrLink && (
								<span className="text-red-500 text-sm">{errors.mrLink}</span>
							)}
						</div>
						<div className="flex flex-col gap-1">
							<Input 
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
						</div>
						<Select
							value={model}
							onChange={(e) => setModel(e.target.value)}
						>
							<SelectItem value="openai">OpenAI</SelectItem>
							<SelectItem value="ollama">Ollama</SelectItem>
						</Select>
					</div>
					<Button 
						className={`w-full sm:w-auto sm:flex-1 flex justify-center items-center  ${
							loading ? "bg-transparent/40 hover:bg-transparent/40 border-none" : ""
						}`} 
						disabled={loading} 
						onClick={handleGenerate}
					>
						{loading ? <HashLoader color="hsl(169,100%,20%)" /> : "Generate"}
					</Button>
				</div>
			</Card>

			<div></div>

		</div>
	)
}

