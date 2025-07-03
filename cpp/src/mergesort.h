
#pragma once

#include <vector>
#include <functional>

namespace
{
    template <typename T>
    void merge(std::vector<T> &arr, int left, int mid, int right, std::function<bool(const T &, const T &)> comparator)
    {
        std::vector<T> tmp;
        tmp.assign(right - left + 1, 0);

        int i = left;
        int j = mid + 1;
        int k = 0;

        while (i <= mid && j <= right)
        {
            if (comparator(arr[i], arr[j]))
            {
                tmp[k++] = arr[i++];
            }
            else
            {
                tmp[k++] = arr[j++];
            }
        }

        for (int idx = i; idx <= mid; ++idx)
        {
            tmp[k++] = arr[idx];
        }

        for (int idx = j; idx <= right; ++idx)
        {
            tmp[k++] = arr[idx];
        }

        for (int idx = 0; idx < tmp.size(); ++idx)
        {
            arr[left + idx] = tmp[idx];
        }
    }

    template <typename T>
    void mergeSort(std::vector<T> &arr, int left, int right, std::function<bool(const T &, const T &)> comparator)
    {
        if (left >= right)
        {
            return;
        }

        int mid = left + ((right - left) / 2);
        mergeSort(arr, left, mid, comparator);
        mergeSort(arr, mid + 1, right, comparator);
        merge(arr, left, mid, right, comparator);
    }
}

namespace sort
{
    template <typename T>
    void MergeSort(std::vector<T> &arr, std::function<bool(const T &, const T &)> comparator = std::less<T>())
    {
        if (arr.size() <= 1)
        {
            return;
        }
        mergeSort(arr, 0, arr.size() - 1, comparator);
    }
}
