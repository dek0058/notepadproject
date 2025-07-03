// See https://aka.ms/new-console-template for more information

using System.Collections.Generic;

List<int> arr = new List<int>() { 1, 5, 2, 8, 3, 9, 4, 7, 6, 10 };

Console.WriteLine("Before Sorting:");
foreach (var val in arr)
{
    Console.Write(val + " ");
}
Console.WriteLine();

//QuickSort<int>.Execute(arr, (x, y) => x > y);
//MergeSort<int>.Execute(arr, (x, y) => x > y);
MinHeapSort<int>.Execute(arr);

Console.WriteLine("After Sorting:");
foreach (var val in arr)
{
    Console.Write(val + " ");
}
Console.WriteLine();
