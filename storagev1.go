package main

import (
    "log"

    storagev1 "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/typed/storage/v1"
)

func applyStorageClass(storageClass *storagev1.StorageClass, storageClassSet v1.StorageClassInterface) error {
    storageClassName := storageClass.GetObjectMeta().GetName()
	storageClasses, err := storageClassSet.List(metav1.ListOptions{})
	if err != nil {
		log.Println("Error when listing storageClass")
		return err
	}

    update := false
    for _, sc := range storageClasses.Items {
        if sc.GetObjectMeta().GetName() == storageClassName {
            update = true
        }
    }

    if update {
        _, err := storageClassSet.Get(storageClassName, metav1.GetOptions{})
        if err != nil {
            log.Println("Error when get old storageClass")
            return err
        }

        _, err = storageClassSet.Update(storageClass)
        if err != nil {
            log.Println("Error when updating storageClass")
            return err
        }
        log.Println("StorageClass " + storageClassName + " updated")

        return err
    } else {
        _, err := storageClassSet.Create(storageClass)
        if err != nil {
            log.Println("Error when creating storageClass")
            return err
        }
        log.Println("StorageClass " + storageClassName + " created")
        return err
    }
}

