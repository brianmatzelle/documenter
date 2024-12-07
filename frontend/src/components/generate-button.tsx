import Button from "@/components/ui/button";
import { HashLoader } from "react-spinners";
import { gitlabColors } from "@/lib/colors";

interface GenerateButtonProps {
  loading: boolean;
  onClick: () => void;
}

export default function GenerateButton({ loading, onClick }: GenerateButtonProps) {
  return (
    <Button 
      className={`w-full p-2 sm:w-auto sm:flex-1 flex justify-center items-center ${
        loading ? "bg-transparent/40 hover:bg-transparent/40 border-none" : ""
      }`} 
      disabled={loading} 
      onClick={onClick}
    >
      {loading ? <HashLoader color={gitlabColors.CINNABAR} /> : "Generate"}
    </Button>
  );
}
