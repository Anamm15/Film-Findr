import { useEffect, useState, useRef } from "react";
import Button from "../../../components/Button";
// eslint-disable-next-line no-unused-vars
import { motion, AnimatePresence } from "framer-motion";
import { updateUser } from "../../../service/user";
import { User, AtSign, Edit3, Save, X, Star, Bookmark, AlignLeft } from "lucide-react";

const ProfileField = ({ label, icon: Icon, value, onChange, disabled }) => {
   return (
      <div className={`group transition-all duration-300 ${disabled ? "" : "bg-gray-50 p-2 rounded-lg border border-gray-200"}`}>
         <label className="text-xs font-semibold text-gray-400 uppercase tracking-wider flex items-center gap-1 mb-1">
            {Icon && <Icon size={14} />} {label}
         </label>
         <input
            disabled={disabled}
            type="text"
            value={value || ""}
            onChange={onChange}
            className={`w-full bg-transparent outline-none text-lg font-medium transition-colors ${disabled ? "text-gray-800" : "text-indigo-700"
               }`}
         />
         {!disabled && <div className="h-0.5 w-0 bg-indigo-500 group-focus-within:w-full transition-all duration-300 rounded-full mt-1" />}
      </div>
   );
};

const Akun = (props) => {
   const { user, review, watchlists } = props;
   const [isUpdating, setIsUpdating] = useState(false);

   // Form States
   const [nama, setNama] = useState(user?.nama || "");
   const [username, setUsername] = useState(user?.username || "");
   const [bio, setBio] = useState(user?.bio || "");

   // UI States
   const [message, setMessage] = useState("");
   const [colorMessage, setColorMessage] = useState("");
   const textareaRef = useRef(null);

   useEffect(() => {
      setNama(user?.nama);
      setUsername(user?.username);
      setBio(user?.bio);
   }, [user]);

   const autoResize = (element) => {
      if (element) {
         element.style.height = "auto";
         element.style.height = element.scrollHeight + "px";
      }
   };

   useEffect(() => {
      autoResize(textareaRef.current);
   }, [bio]);

   const handleUpdateProfile = async () => {
      const data = { nama, username, bio };

      if (user?.id) {
         try {
            const response = await updateUser(user.id, data);
            setIsUpdating(false);
            setMessage(response.data.message);
            setColorMessage("text-green-600 bg-green-50 border-green-200");
            setTimeout(() => setMessage(""), 3000);
         } catch (error) {
            setMessage(error.response?.data?.error || "Update failed");
            setColorMessage("text-red-600 bg-red-50 border-red-200");
         }
      } else {
         alert("User ID not found");
      }
   };

   return (
      <div className="mt-24 mb-10 flex justify-center px-4">
         <div className="w-full max-w-4xl bg-white dark:bg-gray-800 rounded-3xl shadow-xl overflow-hidden border border-gray-100 dark:border-gray-700">

            {/* Header Banner */}
            <div className="relative h-48 bg-gradient-to-r from-indigo-600 to-purple-600">
               <div className="absolute -bottom-16 left-8 md:left-12 flex items-end">
                  <div className="w-32 h-32 rounded-full border-4 border-white dark:border-gray-800 bg-white shadow-lg overflow-hidden flex items-center justify-center">
                     {user?.photo_profil ? (
                        <img
                           src={user.photo_profil}
                           alt={user?.nama || "User"}
                           className="w-full h-full object-cover"
                        />
                     ) : (
                        <span className="text-4xl font-bold text-indigo-600">
                           {user?.nama?.charAt(0).toUpperCase() || "U"}
                        </span>
                     )}
                  </div>
               </div>
            </div>


            <div className="pt-20 px-8 md:px-12 pb-8">
               {/* Header Text & Stats */}
               <div className="flex flex-col md:flex-row justify-between items-start md:items-center gap-6 mb-10">
                  <div>
                     <h1 className="text-3xl font-bold text-gray-900 dark:text-white">
                        {user?.nama || "User Name"}
                     </h1>
                     <p className="text-gray-500 dark:text-gray-400 font-medium">
                        @{user?.username || "username"}
                     </p>
                  </div>

                  <div className="flex gap-4 w-full md:w-auto">
                     <div className="flex-1 md:flex-none flex items-center gap-3 px-5 py-3 bg-blue-50 dark:bg-gray-700/50 rounded-2xl border border-blue-100 dark:border-gray-600">
                        <div className="p-2 bg-blue-100 dark:bg-blue-900/30 text-blue-600 rounded-lg">
                           <Star size={20} />
                        </div>
                        <div>
                           <p className="text-2xl font-bold text-gray-800 dark:text-white leading-none">
                              {review?.reviews?.length || 0}
                           </p>
                           <p className="text-xs text-gray-500 dark:text-gray-400 font-medium uppercase">Reviews</p>
                        </div>
                     </div>

                     <div className="flex-1 md:flex-none flex items-center gap-3 px-5 py-3 bg-purple-50 dark:bg-gray-700/50 rounded-2xl border border-purple-100 dark:border-gray-600">
                        <div className="p-2 bg-purple-100 dark:bg-purple-900/30 text-purple-600 rounded-lg">
                           <Bookmark size={20} />
                        </div>
                        <div>
                           <p className="text-2xl font-bold text-gray-800 dark:text-white leading-none">
                              {watchlists?.length || 0}
                           </p>
                           <p className="text-xs text-gray-500 dark:text-gray-400 font-medium uppercase">Watchlist</p>
                        </div>
                     </div>
                  </div>
               </div>

               <hr className="border-gray-100 dark:border-gray-700 mb-8" />

               {/* Inputs */}
               <div className="grid grid-cols-1 md:grid-cols-2 gap-8 mb-8">
                  <ProfileField
                     label="Full Name"
                     icon={User}
                     value={nama}
                     onChange={(e) => setNama(e.target.value)}
                     disabled={!isUpdating}
                  />
                  <ProfileField
                     label="Username"
                     icon={AtSign}
                     value={username}
                     onChange={(e) => setUsername(e.target.value)}
                     disabled={!isUpdating}
                  />
               </div>

               {/* Bio */}
               <div className={`mb-8 transition-all duration-300 ${!isUpdating ? "" : "bg-gray-50 p-4 rounded-xl border border-gray-200"}`}>
                  <label className="text-xs font-semibold text-gray-400 uppercase tracking-wider flex items-center gap-1 mb-2">
                     <AlignLeft size={14} /> About Me
                  </label>
                  <textarea
                     ref={textareaRef}
                     disabled={!isUpdating}
                     value={bio}
                     onChange={(e) => {
                        setBio(e.target.value);
                        autoResize(e.target);
                     }}
                     className={`w-full overflow-hidden resize-none bg-transparent outline-none leading-relaxed transition-colors ${!isUpdating
                        ? "text-gray-600 dark:text-gray-300 border-none p-0"
                        : "text-gray-800 min-h-[100px]"
                        }`}
                     placeholder={isUpdating ? "Tell us something about yourself..." : "No bio yet."}
                  />
               </div>

               {/* Buttons & AnimatePresence */}
               <div className="flex flex-col md:flex-row justify-between items-center gap-4 mt-10">
                  <div className="order-2 md:order-1 flex-1">
                     <AnimatePresence>
                        {message && (
                           // --- Menggunakan motion.div di sini ---
                           <motion.div
                              initial={{ opacity: 0, x: -10 }}
                              animate={{ opacity: 1, x: 0 }}
                              exit={{ opacity: 0 }}
                              className={`px-4 py-2 rounded-lg text-sm font-medium border ${colorMessage} w-max`}
                           >
                              {message}
                           </motion.div>
                        )}
                     </AnimatePresence>
                  </div>

                  <div className="order-1 md:order-2 flex gap-3 w-full md:w-auto justify-end">
                     <AnimatePresence mode="wait">
                        {isUpdating ? (
                           // --- Menggunakan motion.div di sini ---
                           <motion.div
                              key="actions"
                              initial={{ opacity: 0, y: 10 }}
                              animate={{ opacity: 1, y: 0 }}
                              exit={{ opacity: 0, y: 10 }}
                              className="flex gap-3 w-full md:w-auto"
                           >
                              <Button
                                 variant="outline"
                                 onClick={() => {
                                    setIsUpdating(false);
                                    setNama(user?.nama);
                                    setUsername(user?.username);
                                    setBio(user?.bio);
                                 }}
                                 className="flex items-center justify-center gap-2 rounded-xl px-6 py-2.5 border-gray-300 hover:bg-gray-50 text-gray-700 w-1/2 md:w-auto"
                              >
                                 <X size={18} /> Cancel
                              </Button>
                              <Button
                                 onClick={handleUpdateProfile}
                                 className="flex items-center justify-center gap-2 bg-indigo-600 hover:bg-indigo-700 text-white rounded-xl px-6 py-2.5 shadow-lg shadow-indigo-200 w-1/2 md:w-auto"
                              >
                                 <Save size={18} /> Save Changes
                              </Button>
                           </motion.div>
                        ) : (
                           // --- Menggunakan motion.div di sini ---
                           <motion.div
                              key="edit"
                              initial={{ opacity: 0, scale: 0.9 }}
                              animate={{ opacity: 1, scale: 1 }}
                              exit={{ opacity: 0, scale: 0.9 }}
                           >
                              <Button
                                 onClick={() => setIsUpdating(true)}
                                 className="flex items-center gap-2 bg-gray-900 dark:bg-gray-700 hover:bg-black text-white rounded-xl px-6 py-2.5 shadow-lg transition-transform active:scale-95"
                              >
                                 <Edit3 size={18} /> Edit Profile
                              </Button>
                           </motion.div>
                        )}
                     </AnimatePresence>
                  </div>
               </div>

            </div>
         </div>
      </div>
   );
};

export default Akun;