// components/EditCategoryModal.tsx
"use client";

import { useState, useEffect } from "react";
import Snackbar from "@/components/Snackbar"; // Assuming Snackbar is in components

interface EditCategoryModalProps {
  isOpen: boolean;
  onClose: () => void;
  onEdit: (categoryId: string, newCategoryName: string) => Promise<void>;
  editCategoryError: string | null;
  isEditingCategory: boolean;
  handleCloseSnackbar: () => void;
  categoryName: string; // Initial category name
  categoryId: string;
}

const EditCategoryModal: React.FC<EditCategoryModalProps> = ({
  isOpen,
  onClose,
  onEdit,
  editCategoryError,
  isEditingCategory,
  handleCloseSnackbar,
  categoryName: initialCategoryName,
  categoryId,
}) => {
  const [editedCategoryName, setEditedCategoryName] = useState(initialCategoryName);

  useEffect(() => {
    if (isOpen) {
      setEditedCategoryName(initialCategoryName);
    }
  }, [isOpen, initialCategoryName]);

  const handleEdit = async () => {
    await onEdit(categoryId, editedCategoryName);
  };

  return (
    <dialog
      id="edit_category_modal"
      className="modal modal-bottom sm:modal-middle backdrop:bg-gray-900 backdrop:bg-opacity-50"
      open={isOpen}
    >
      <div className="modal-box bg-gray-800 text-gray-100">
        <h3 className="font-bold text-lg font-mono">Edit Category</h3>
        <form method="dialog" className="mt-4">
          <div className="mb-4">
            <label
              htmlFor="editCategoryName"
              className="block text-gray-300 text-sm font-bold mb-2 font-mono"
            >
              Category Name
            </label>
            <input
              type="text"
              id="editCategoryName"
              placeholder="Enter new category name"
              className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-100 leading-tight focus:outline-none focus:shadow-outline bg-gray-700 border-gray-600 font-mono"
              value={editedCategoryName}
              onChange={(e) => setEditedCategoryName(e.target.value)}
            />
          </div>
          {editCategoryError && (
            <Snackbar
              type="error"
              message={editCategoryError}
              onClose={handleCloseSnackbar}
            />
          )}
          <div className="modal-action">
            <button className="btn btn-error font-mono" onClick={onClose}>
              Cancel
            </button>
            <button
              className="btn btn-primary font-mono"
              onClick={handleEdit}
              disabled={!editedCategoryName.trim() || isEditingCategory}
            >
              Save
            </button>
          </div>
        </form>
      </div>
    </dialog>
  );
};

export default EditCategoryModal;
