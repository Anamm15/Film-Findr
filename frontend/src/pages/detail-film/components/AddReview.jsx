import { useContext, useState } from "react";
import { Star, Send } from "lucide-react"; 
import { createReview } from "../../../service/review";
import Button from "../../../components/Button";
import TextArea from "../../../components/Textarea";
import { INIT_RATING_REVIEW } from "../../../utils/constant";
import { AuthContext } from "../../../contexts/AuthContext";

const AddReview = (props) => {
    const { id, setReviews } = props;
    const { user } = useContext(AuthContext);

    const [message, setMessage] = useState("");
    const [colorMessage, setColorMessage] = useState("");

    // State untuk form
    const [newReview, setNewReview] = useState("");
    const [rating, setRating] = useState(INIT_RATING_REVIEW);
    const [hoverRating, setHoverRating] = useState(0);

    const handleAddReview = async (e) => {
        e.preventDefault();

        // Validasi sederhana
        if (rating === 0) {
            setMessage("Please select a star rating");
            setColorMessage("text-red-500");
            return;
        }

        try {
            const data = {
                film_id: id,
                komentar: newReview,
                rating: Number(rating)
            };

            const response = await createReview(data);

            setMessage(response.data.message);
            setColorMessage("text-green-600");

            // Reset form
            setNewReview("");
            setRating(0);

            setReviews((prevReviews) => ({
                ...prevReviews,
                reviews: [...prevReviews.reviews, {
                    id: response.data.id,
                    komentar: response.data.komentar,
                    rating: response.data.rating,
                    user: user
                }]
            }));
        } catch (error) {
            setMessage(error.response?.data?.error || "Failed to add review");
            setColorMessage("text-red-600");
            console.log(error);
        }
    };

    return (
        <div className="bg-white dark:bg-gray-800 rounded-2xl shadow-lg border border-gray-100 dark:border-gray-700 p-8 transition-all duration-300 hover:shadow-xl">
            <div className="mb-6">
                <h2 className="text-2xl font-bold text-gray-800 dark:text-white">Review This Movie</h2>
                <p className="text-gray-500 dark:text-gray-400 text-sm mt-1">
                    Share your thoughts with other viewers
                </p>
            </div>

            <form onSubmit={handleAddReview} className="flex flex-col gap-6">
                {/* Section Rating Bintang */}
                <div className="flex flex-col gap-2">
                    <label className="text-sm font-semibold text-gray-700 dark:text-gray-300 uppercase tracking-wider">
                        Your Rating
                    </label>
                    <div className="flex items-center gap-1">
                        {[...Array(10)].map((_, index) => {
                            const starValue = index + 1;
                            return (
                                <button
                                    type="button"
                                    key={starValue}
                                    className="focus:outline-none transition-transform hover:scale-110 active:scale-95"
                                    onClick={() => setRating(starValue)}
                                    onMouseEnter={() => setHoverRating(starValue)}
                                    onMouseLeave={() => setHoverRating(0)}
                                >
                                    <Star
                                        size={28}
                                        className={`${starValue <= (hoverRating || rating)
                                            ? "fill-yellow-400 text-yellow-400"
                                            : "fill-transparent text-gray-300 dark:text-gray-600"
                                            } transition-colors duration-200`}
                                    />
                                </button>
                            );
                        })}
                        <span className="ml-3 text-lg font-bold text-gray-700 dark:text-gray-200 w-8">
                            {hoverRating || rating || 0}
                            <span className="text-xs text-gray-400 font-normal">/10</span>
                        </span>
                    </div>
                </div>

                {/* Section Text Area */}
                <div className="flex flex-col gap-2">
                    <label className="text-sm font-semibold text-gray-700 dark:text-gray-300 uppercase tracking-wider">
                        Your Comment
                    </label>
                    <TextArea
                        placeholder="Write something insightful about the movie..."
                        value={newReview}
                        onChange={(e) => setNewReview(e.target.value)}
                        className="w-full min-h-[120px] p-4 bg-gray-50 dark:bg-gray-900 border border-gray-200 dark:border-gray-700 rounded-xl focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all outline-none resize-y"
                    />
                </div>

                {/* Section Button & Message */}
                <div className="flex flex-col items-end gap-3">
                    <Button
                        type="submit"
                        className="flex items-center gap-2 px-8 py-3 bg-blue-600 hover:bg-blue-700 text-white rounded-xl font-medium transition-all shadow-md hover:shadow-lg disabled:opacity-50 disabled:cursor-not-allowed"
                        disabled={!newReview || rating === 0}
                    >
                        <span>Post Review</span>
                        <Send size={18} />
                    </Button>

                    {message && (
                        <div className={`text-sm font-medium px-4 py-2 rounded-lg bg-opacity-10 ${colorMessage.includes('red') ? 'bg-red-100' : 'bg-green-100'} ${colorMessage}`}>
                            {message}
                        </div>
                    )}
                </div>
            </form>
        </div>
    );
};

export default AddReview;