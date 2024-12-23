// components/AddCategoryModal.tsx
"use client";

import { useState } from "react";
import Snackbar from "@/components/Snackbar"; // Assuming Snackbar is in components

interface AddCategoryModalProps {
  isOpen: boolean;
  onClose: () => void;
  onCreate: (categoryName: string) => Promise<void>;
  createCategoryError: string | null;
  isCreatingCategory: boolean;
  handleCloseSnackbar: () => void;
}

const AddCategoryModal: React.FC<AddCategoryModalProps> = ({
  isOpen,
  onClose,
  onCreate,
  createCategoryError,
  isCreatingCategory,
  handleCloseSnackbar,
}) => {
  const [newCategoryName, setNewCategoryName] = useState("");

  const handleCreate = async () => {
    await onCreate(newCategoryName);
    setNewCategoryName("");
  };

  return (
    <dialog
      id="add_category_modal"
      className="modal modal-bottom sm:modal-middle backdrop:bg-gray-900 backdrop:bg-opacity-50"
      open={isOpen}
    >
      <div className="modal-box bg-gray-800 text-gray-100">
        <h3 className="font-bold text-lg font-mono">Add New Category</h3>
        <form method="dialog" className="mt-4">
          <div className="mb-4">
            <label
              htmlFor="categoryName"
              className="block text-gray-300 text-sm font-bold mb-2 font-mono"
            >
              Category Name
            </label>
            <input
              type="text"
              id="categoryName"
              placeholder="Enter category name"
              className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-100 leading-tight focus:outline-none focus:shadow-outline bg-gray-700 border-gray-600 font-mono"
              value={newCategoryName}
              onChange={(e) => setNewCategoryName(e.target.value)}
            />
          </div>
          {createCategoryError && (
            <Snackbar
              type="error"
              message={createCategoryError}
              onClose={handleCloseSnackbar}
            />
          )}
          <div className="modal-action">
            <button className="btn btn-error font-mono" onClick={onClose}>
              Cancel
            </button>
            <button
              className="btn btn-primary font-mono"
              onClick={handleCreate}
              disabled={!newCategoryName.trim() || isCreatingCategory}
            >
              Create
            </button>
          </div>
        </form>
      </div>
    </dialog>
  );
};

export default AddCategoryModal;
