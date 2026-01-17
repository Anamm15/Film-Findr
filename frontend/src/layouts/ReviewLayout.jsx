import { ThumbsUp, ThumbsDown, Star, MessageSquare } from "lucide-react";
import Pagination from "../components/Pagination";
import { updateReaksiReview } from "../service/review";
import Loading from "../components/Loading";

const ReviewLayout = (props) => {
   const { review, setReviews, setPage, page, loading } = props;

   const handleReact = async (id, reaksi, idx) => {
      try {
         const payload = {
            reaksi: reaksi
         };
         const response = await updateReaksiReview(id, payload);
         if (response) {
            const newReviews = [...review.reviews];
            // Update state reaksi user secara optimis
            newReviews[idx] = {
               ...newReviews[idx],
               user_reaksi: {
                  reaksi: reaksi
               }
            };
            setReviews({ ...review, reviews: newReviews });
         }
      } catch (error) {
         console.log(error);
      }
   };

   // Helper untuk membuat avatar inisial dari nama
   const getInitials = (name) => {
      if (!name) return "U";
      return name
         .split(" ")
         .map((n) => n[0])
         .slice(0, 2)
         .join("")
         .toUpperCase();
   };

   return (
      <div className="bg-white dark:bg-gray-800 rounded-2xl shadow-lg border border-gray-100 dark:border-gray-700 p-6 mb-8">
         <div className="flex items-center gap-3 mb-6 ps-2">
            <MessageSquare className="text-blue-600" />
            <h2 className="text-2xl font-bold text-gray-800 dark:text-white">
               User Reviews <span className="text-gray-400 text-lg font-normal">({review?.total || 0})</span>
            </h2>
         </div>

         <div className="space-y-6">
            {loading ? (
               <div className="py-10 flex justify-center">
                  <Loading>Loading reviews...</Loading>
               </div>
            ) : review && review.reviews && review.reviews.length > 0 ? (
               review.reviews.map((r, idx) => (
                  <div
                     key={r.id || idx}
                     className="flex flex-col gap-3 pb-6 border-b border-gray-100 dark:border-gray-700 last:border-0 last:pb-0"
                  >
                     {/* Header Review: Avatar, Nama, Rating */}
                     <div className="flex justify-between items-start">
                        <div className="flex items-center gap-3">
                           {/* Avatar Circle */}
                           <div className="w-10 h-10 rounded-full bg-gradient-to-br from-blue-500 to-purple-600 flex items-center justify-center text-white font-bold text-sm shadow-sm">
                              {
                                 r.user?.photo_profil
                                    ? (
                                       <img
                                          className="w-full h-full rounded-full object-cover"
                                          src={r.user?.photo_profil}
                                          alt={r.user?.username}
                                       />
                                    )
                                    : (
                                       <span>{getInitials(r.user?.username)}</span>
                                    )
                              }
                           </div>

                           <div className="flex flex-col">
                              <p className="font-semibold text-gray-900 dark:text-gray-100 leading-tight">
                                 {r.user?.username || "Anonymous"}
                              </p>
                              {/* Tampilan Rating */}
                              <div className="flex items-center gap-1 mt-1">
                                 <Star size={14} className="fill-yellow-400 text-yellow-400" />
                                 <span className="text-sm font-bold text-gray-700 dark:text-gray-300">
                                    {r.rating}
                                 </span>
                                 <span className="text-xs text-gray-400">/ 10</span>
                              </div>
                           </div>
                        </div>

                        {/* Tanggal (Opsional, jika ada data created_at) */}
                        {/* <span className="text-xs text-gray-400">2 days ago</span> */}
                     </div>

                     {/* Isi Komentar */}
                     <p className="text-gray-600 dark:text-gray-300 leading-relaxed text-sm md:text-base pl-[52px]">
                        {r.komentar}
                     </p>

                     {/* Action Buttons (Like/Dislike) */}
                     <div className="flex items-center gap-4 pl-[52px] mt-1">
                        <button
                           onClick={() => handleReact(r.id, "like", idx)}
                           className={`group flex items-center gap-1.5 text-sm transition-colors duration-200 ${r.user_reaksi?.reaksi === "like"
                              ? "text-blue-600 font-medium"
                              : "text-gray-400 hover:text-blue-500"
                              }`}
                        >
                           <ThumbsUp
                              size={16}
                              className={`transition-transform group-active:scale-125 ${r.user_reaksi?.reaksi === "like" ? "fill-blue-600" : ""
                                 }`}
                           />
                           <span>Helpful</span>
                        </button>

                        <button
                           onClick={() => handleReact(r.id, "dislike", idx)}
                           className={`group flex items-center gap-1.5 text-sm transition-colors duration-200 ${r.user_reaksi?.reaksi === "dislike"
                              ? "text-red-500 font-medium"
                              : "text-gray-400 hover:text-red-500"
                              }`}
                        >
                           <ThumbsDown
                              size={16}
                              className={`transition-transform group-active:scale-125 mt-0.5 ${r.user_reaksi?.reaksi === "dislike" ? "fill-red-500" : ""
                                 }`}
                           />
                           <span>Not helpful</span>
                        </button>
                     </div>
                  </div>
               ))
            ) : (
               <div className="text-center py-10 text-gray-400">
                  <MessageSquare size={48} className="mx-auto mb-3 opacity-20" />
                  <p>No reviews yet. Be the first to share your thoughts!</p>
               </div>
            )}
         </div>

         {/* Pagination hanya muncul jika ada review */}
         {review && review.reviews && review.reviews.length > 0 && (
            <div className="mt-8 pt-4 border-t border-gray-100 dark:border-gray-700">
               <Pagination contents={review} page={page} setPage={setPage} />
            </div>
         )}
      </div>
   );
};

export default ReviewLayout;