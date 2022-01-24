# Fluxtomizer

1. Run kustomize for environment
2. Go through yaml documents, check patches have targets defined and create an object with base path and patches
3. for object in list of objects, update the kustomization.yaml at the base path and run kustomize. if errors are present store the errors.
4. Print successful result or kustomize outputs or errors


Need:
- [X] function to run kustomize
- [] function to check patches
- [] function to loop over yaml documents, run check function and create new object
- [] function to update kustomization.yaml with patches
- [] function to remove patches from kustomization.yaml
- [] function to loop objects, update k.yaml, run kustomize, store errors or output