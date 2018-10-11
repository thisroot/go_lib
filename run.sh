#check changes on a shared package library
    if ! git diff-index --quiet HEAD --; then
        echo "Changes detected into the shared go_lib package library"
        git add .
        git commit -m $(date +%Y%m%d%H%M%S)
        git push origin master
        exit 1
    fi

    echo "Changes into the shared go_lib package library undetected"